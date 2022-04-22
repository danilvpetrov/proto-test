package pong

import (
	"time"

	"github.com/danilvpetrov/proto-test/data"
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
	req *data.PingRequest,
	stream data.PingPong_DoPingPongServer,
) error {
	if req.Ping.Text != "Ping" {
		return status.Error(codes.InvalidArgument, "Invalid argument in the request")
	}

	t := time.NewTicker(2 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-stream.Context().Done():
			return status.Error(codes.Aborted, stream.Context().Err().Error())
		case <-t.C:
			if err := stream.Send(
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
