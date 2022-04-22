package main

import (
	"log"
	"net"
	"os"

	"github.com/danilvpetrov/proto-test/data"
	"github.com/danilvpetrov/proto-test/pong"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PONG_SERVER_PORT")
	if port == "" {
		panic("environment variable PONG_SERVER_PORT is empty or not defined")
	}

	addr := net.JoinHostPort("", port)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	log.Printf("pong server is listening on %v", l.Addr())

	svr := grpc.NewServer()
	data.RegisterPingPongServer(svr, &pong.Server{})

	if err := svr.Serve(l); err != nil {
		panic(err)
	}
}
