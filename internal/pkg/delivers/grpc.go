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
	server   *grpc.Server
}

func NewGRPCDeliver() *GRPCDeliver {
	return &GRPCDeliver{
		listener: newTCPListener(),
		server:   grpc.NewServer(),
	}
}

func (d *GRPCDeliver) Server() *grpc.Server {
	return d.server
}

func (d *GRPCDeliver) Run() {
	reflection.Register(d.server)
	if err := d.server.Serve(d.listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func newTCPListener() net.Listener {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return listener
}
