package delivers

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

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

func lookupTCPPort() uint16 {
	strHost, ok := os.LookupEnv("PORT")
	if !ok {
		return uint16(50051)
	}
	host, err := strconv.ParseUint(strHost, 10, 16)
	if err != nil {
		panic(err)
	}
	return uint16(host)
}

func newTCPListener() net.Listener {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", lookupTCPPort()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return listener
}
