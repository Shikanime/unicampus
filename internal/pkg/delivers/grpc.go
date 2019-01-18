package delivers

import (
	"log"
	"net"

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
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return listener
}
