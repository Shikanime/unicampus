package main

import (
	"github.com/Shikanime/unicampus/cmd/admission/app"
	"github.com/Shikanime/unicampus/cmd/admission/delivers/grpc"
	"github.com/Shikanime/unicampus/cmd/admission/services/indexer"
	"github.com/Shikanime/unicampus/cmd/admission/services/persistence"
	"github.com/Shikanime/unicampus/pkg/unicampus_api_admission_v1"
)

type Server struct {
	app.School
	app.Student
}

func main() {
	grpcDeliver := grpc.NewServer()

	persistenceRepo := persistence.NewClient()
	defer persistenceRepo.Close()
	indexerRepo := indexer.NewClient()
	defer indexerRepo.Close()

	unicampus_api_admission_v1.RegisterAdmissionServiceServer(grpcDeliver.Driver(), &Server{
		School: app.NewSchool(&persistenceRepo, &indexerRepo),
		// Student: app.NewStudent(&persistenceRepo, &indexerRepo),
	})

	grpcDeliver.Run()
}
