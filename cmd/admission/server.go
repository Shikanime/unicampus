package main

import (
	"github.com/Shikanime/unicampus/api/admission/v1alpha1"
	"github.com/Shikanime/unicampus/cmd/admission/app"
	"github.com/Shikanime/unicampus/cmd/admission/delivers"
	"github.com/Shikanime/unicampus/cmd/admission/repositories/indexer"
	"github.com/Shikanime/unicampus/cmd/admission/repositories/persistence"
	"github.com/Shikanime/unicampus/cmd/admission/services"
)

type Server struct {
	app.School
	app.Student
}

func main() {
	grpcDeliver := delivers.NewGRPCDeliver()
	defer grpcDeliver.Run()

	postgresService := services.NewPostgresService()
	defer postgresService.Close()
	elasticserachService := services.NewElasticSearchService()
	defer elasticserachService.Close()

	persistenceRepo := persistence.NewRepository(postgresService)
	indexerRepo := indexer.NewRepository(elasticserachService)

	unicampus_api_admission_v1alpha1.RegisterAdmissionServiceServer(grpcDeliver.Driver(), &Server{
		School: app.NewSchool(&persistenceRepo, &indexerRepo),
		// Student: app.NewStudent(&persistenceRepo, &indexerRepo),
	})
}
