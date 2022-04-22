package main

import (
	"context"
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

	client := data.NewPingPongClient(conn)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	res, err := client.DoPingPong(ctx, &data.PingRequest{
		Ping: &data.Ping{
			Text: "Ping",
		},
	})
	if err != nil {
		panic(err)
	}

	log.Printf("Response recieved: %v", res)

}
