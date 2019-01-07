package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Shikanime/unicampus/cmd/admission/containers"
	"github.com/Shikanime/unicampus/cmd/admission/services/indexer"
	"github.com/Shikanime/unicampus/cmd/admission/services/persistence"
	"github.com/Shikanime/unicampus/pkg/unicampus_api_admission_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	containers.School
	containers.Student
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

	persistenceRepo := persistence.NewClient()
	defer persistenceRepo.Close()
	indexerRepo := indexer.NewClient()
	defer indexerRepo.Close()

	// Server
	s := grpc.NewServer()
	unicampus_api_admission_v1.RegisterAdmissionServiceServer(s, &Server{
		School: containers.NewSchool(&persistenceRepo, &indexerRepo),
		// Student: containers.NewStudent(&persistenceRepo, &indexerRepo),
	})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(tcpListener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
