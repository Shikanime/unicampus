package main

import (
	unicampus_api_admission_v1alpha1 "github.com/Shikanime/unicampus/api/admission/v1alpha1"
	app "github.com/Shikanime/unicampus/internal/app/admission"
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/indexer"
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/persistence"
	"github.com/Shikanime/unicampus/internal/pkg/delivers"
	"github.com/Shikanime/unicampus/internal/pkg/services"
)

type App struct {
	app.School
	app.Student
	app.Application
}

func main() {
	grpcDeliver := delivers.NewGRPCDeliver()

	postgresService := services.NewPostgresService()
	defer postgresService.Close()
	elasticserachService := services.NewElasticSearchService()
	defer elasticserachService.Close()

	persistenceRepo := persistence.NewRepository(postgresService)
	indexerRepo := indexer.NewRepository(elasticserachService)

	unicampus_api_admission_v1alpha1.RegisterAdmissionServiceServer(grpcDeliver.Server(), &App{
		School:      app.NewSchoolService(&persistenceRepo, &indexerRepo),
		Student:     app.NewStudentService(&persistenceRepo, &indexerRepo),
		Application: app.NewApplicationService(&persistenceRepo, &indexerRepo),
	})

	grpcDeliver.Run()
}
