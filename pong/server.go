package pong

import (
	"context"

	"github.com/danilvpetrov/proto-test/data"
)

// Server is an implementation of data.PingPongServer.
type Server struct {
	data.UnimplementedPingPongServer
}

// DoPingPong performs handling of the Ping request from the client and returns
// Pong response.
func (s *Server) DoPingPong(
	ctx context.Context,
	req *data.PingRequest,
) (*data.PongResponse, error) {
	return &data.PongResponse{
		Pong: &data.Pong{
			Text: "Pong",
		},
	}, nil
}
