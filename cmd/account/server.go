package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Shikanime/unicampus/cmd/account/persistence"
	"github.com/Shikanime/unicampus/pkg/account"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	persistence *persistence.Repo
}

func (s *Server) SignIn(context.Context, *unicampus_account.SignInRequest) (*unicampus_account.SignInReply, error) {
	return &unicampus_account.SignInReply{}, nil
}

func (s *Server) SignUp(context.Context, *unicampus_account.SignUpRequest) (*unicampus_account.SignUpReply, error) {
	return &unicampus_account.SignUpReply{}, nil

}

func (s *Server) SignOut(context.Context, *unicampus_account.SignOutRequest) (*unicampus_account.SignOutReply, error) {
	return &unicampus_account.SignOutReply{}, nil
}

func (s *Server) Close(context.Context, *unicampus_account.CloseRequest) (*unicampus_account.CloseReply, error) {
	return &unicampus_account.CloseReply{}, nil

}

func NewTCPListener() net.Listener {
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

func main() {
	tcpListener := NewTCPListener()

	// Server
	s := grpc.NewServer()
	unicampus_account.RegisterAccountServiceServer(s, &Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(tcpListener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
