package delivers

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCDeliver struct {
	listener net.Listener
	driver   *grpc.Server
}

func NewGRPCDeliver() *GRPCDeliver {
	return &GRPCDeliver{
		listener: newTCPListener(),
		driver:   grpc.NewServer(),
	}
}

func (s *GRPCDeliver) Driver() *grpc.Server {
	return s.driver
}

func (s *GRPCDeliver) Run() {
	reflection.Register(s.driver)
	if err := s.driver.Serve(s.listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func newTCPListener() net.Listener {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return listener
}
