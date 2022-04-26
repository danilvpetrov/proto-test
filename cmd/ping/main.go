package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/danilvpetrov/proto-test/data"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	addr := os.Getenv("PONG_SERVER_ADDR")
	if addr == "" {
		panic("environment variable PONG_SERVER_ADDR is empty or not defined")
	}

	creds := grpc.WithTransportCredentials(insecure.NewCredentials())

	conn, err := grpc.Dial(addr, creds)
	if err != nil {
		panic(err)
	}
	defer conn.Close() //nolint

	client := data.NewPingPongClient(conn)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	stream, err := client.DoPingPong(ctx)
	if err != nil {
		panic(err)
	}

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return recv(ctx, stream)
	})

	eg.Go(func() error {
		return send(ctx, stream)
	})

	eg.Wait()

}

func recv(ctx context.Context, receiver interface {
	Recv() (*data.PongResponse, error)
}) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		resp, err := receiver.Recv()
		if err != nil {
			return err
		}

		if resp.Pong.Text != "Pong" {
			return status.Error(codes.InvalidArgument, "Invalid argument in the response")
		}

		log.Printf("received a response %v", resp.Pong)
	}
}

func send(ctx context.Context, sender interface {
	Send(*data.PingRequest) error
}) error {
	t := time.NewTicker(2 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return status.Error(codes.Aborted, ctx.Err().Error())
		case <-t.C:
			if err := sender.Send(
				&data.PingRequest{
					Ping: &data.Ping{
						Text: "Ping",
					},
				},
			); err != nil {
				return status.Error(codes.Internal, err.Error())
			}
		}
	}
}
