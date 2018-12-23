package server

import (
	"context"

	hub "github.com/shellimsi/proto/hub"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	Port string
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Resize(ctx context.Context, in *hub.WinsizeRequest) (*hub.WinsizeResponse, error) {
	ok := status("OK")
	return &hub.WinsizeResponse{Status: ok}, nil
}
