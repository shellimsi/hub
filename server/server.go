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
	return &hub.WinsizeResponse{Status: hub.Status_OK}, nil
}

func (s *Server) Register(ctx context.Context, register *hub.RegisterRequest) (*hub.RegisterResponse, error) {
	return &hub.RegisterResponse{
		Host:      "localhost",
		Port:      8080,
		SessionId: 1,
	}, nil
}
