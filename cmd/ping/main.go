package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/danilvpetrov/proto-test/data"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	stream, err := client.DoPingPong(ctx, &data.PingRequest{
		Ping: &data.Ping{
			Text: "Ping",
		},
	})
	if err != nil {
		panic(err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Print("stream closed")
			return
		}
		if err != nil {
			log.Fatalf("cannot receive %v", err)
		}
		log.Printf("resp received: %s", resp.Pong.Text)
	}

}
