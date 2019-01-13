package main

import (
	unicampus_api_admission_v1alpha1 "github.com/Shikanime/unicampus/api/admission/v1alpha1"
	app "github.com/Shikanime/unicampus/internal/app/admission"
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/indexer"
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/persistence"
	"github.com/Shikanime/unicampus/internal/pkg/delivers"
	"github.com/Shikanime/unicampus/internal/pkg/services"
	"github.com/Shikanime/unicampus/pkg/admission"
)

type App struct {
	app.School
	app.Student
	app.Application
}

func main() {
	grpcDeliver := delivers.NewGRPCDeliver()

	postgresService := services.NewPostgreSQLService()
	defer postgresService.Close()
	elasticserachService := services.NewElasticSearchService()
	defer elasticserachService.Close()

	persistenceRepo := persistence.NewRepository(postgresService)
	indexerRepo := indexer.NewRepository(elasticserachService)

	schoolService := app.NewSchoolService(&persistenceRepo, &indexerRepo)
	studentService := app.NewStudentService(&persistenceRepo, &indexerRepo)
	applicationService := app.NewApplicationService(&persistenceRepo, &indexerRepo)

	var err error
	if err = persistenceRepo.Init(); err != nil {
		panic(err)
	}
	if err = indexerRepo.Init(); err != nil {
		panic(err)
	}

	persistenceRepo.CreateSchool(&admission.School{
		UUID:        "yo",
		Name:        "ETNA",
		Description: "Desc",
	})

	unicampus_api_admission_v1alpha1.RegisterAdmissionServiceServer(grpcDeliver.Server(), &App{
		School:      schoolService,
		Student:     studentService,
		Application: applicationService,
	})

	grpcDeliver.Run()
}
