package server

import (
	"net"

	hub "github.com/shellimsi/proto/hub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func status(status string) string {
	v, ok := hub.Status_value[status]
	if ok {
		panic("invalid enum value")
	}
	vStr, ok := hub.Status_name[v]
	if ok {
		panic("invalid enum value")
	}
	return vStr
}

func Serve(service *Server) error {
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	hub.RegisterTerminalServer(s, &Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
