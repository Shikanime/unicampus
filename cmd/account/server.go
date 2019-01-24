package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	unicampus_api_account_v1alpha1 "gitlab.com/deva-hub/unicampus/api/account/v1alpha1"
	"gitlab.com/deva-hub/unicampus/internal/app/admission/repositories/persistence"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	persistence *persistence.Repo
}

func (s *Server) SignIn(context.Context, *unicampus_api_account_v1alpha1.SignInRequest) (*unicampus_api_account_v1alpha1.SignInReply, error) {
	return &unicampus_api_account_v1alpha1.SignInReply{}, nil
}

func (s *Server) SignUp(context.Context, *unicampus_api_account_v1alpha1.SignUpRequest) (*unicampus_api_account_v1alpha1.SignUpReply, error) {
	return &unicampus_api_account_v1alpha1.SignUpReply{}, nil
}

func (s *Server) SignOut(context.Context, *unicampus_api_account_v1alpha1.SignOutRequest) (*unicampus_api_account_v1alpha1.SignOutReply, error) {
	return &unicampus_api_account_v1alpha1.SignOutReply{}, nil
}

func (s *Server) Close(context.Context, *unicampus_api_account_v1alpha1.CloseRequest) (*unicampus_api_account_v1alpha1.CloseReply, error) {
	return &unicampus_api_account_v1alpha1.CloseReply{}, nil
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
	unicampus_api_account_v1alpha1.RegisterAccountServiceServer(s, &Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(tcpListener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
