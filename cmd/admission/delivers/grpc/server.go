package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPC struct {
	listener net.Listener
	driver   *grpc.Server
}

func NewServer() *GRPC {
	return &GRPC{
		listener: newTCPListener(),
		driver:   grpc.NewServer(),
	}
}

func (s *GRPC) Driver() *grpc.Server {
	return s.driver
}

func (s *GRPC) Run() {
	reflection.Register(s.driver)
	if err := s.driver.Serve(s.listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
