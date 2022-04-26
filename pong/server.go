package pong

import (
	"context"
	"log"
	"time"

	"github.com/danilvpetrov/proto-test/data"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server is an implementation of data.PingPongServer.
type Server struct {
	data.UnimplementedPingPongServer
}

// DoPingPong performs handling of the Ping request from the client and returns
// Pong response.
func (s *Server) DoPingPong(
	stream data.PingPong_DoPingPongServer,
) error {
	ctx, cancel := context.WithCancel(stream.Context())
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return recv(ctx, stream)
	})

	eg.Go(func() error {
		return send(ctx, stream)
	})

	return eg.Wait()
}

func recv(ctx context.Context, receiver interface {
	Recv() (*data.PingRequest, error)
}) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		req, err := receiver.Recv()
		if err != nil {
			return err
		}

		if req.Ping.Text != "Ping" {
			return status.Error(codes.InvalidArgument, "Invalid argument in the request")
		}

		log.Printf("received a request %v", req)
	}
}

func send(ctx context.Context, sender interface {
	Send(*data.PongResponse) error
}) error {
	t := time.NewTicker(2 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return status.Error(codes.Aborted, ctx.Err().Error())
		case <-t.C:
			if err := sender.Send(
				&data.PongResponse{
					Pong: &data.Pong{
						Text: "Pong",
					},
				},
			); err != nil {
				return status.Error(codes.Internal, err.Error())
			}
		}
	}
}
