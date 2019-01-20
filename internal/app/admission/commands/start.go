package commands

import (
	unicampus_api_admission_v1alpha1 "github.com/Shikanime/unicampus/api/admission/v1alpha1"
	app "github.com/Shikanime/unicampus/internal/app/admission"
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/indexer"
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/persistence"
	"github.com/Shikanime/unicampus/internal/pkg/delivers"
	"github.com/Shikanime/unicampus/internal/pkg/services"
	"github.com/spf13/cobra"
)

type server struct {
	app.School
	app.Student
	app.Application
}

func NewStart(appName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start service",
		Long:  `start command is used to start the service.`,
		Run: func(cmd *cobra.Command, args []string) {
			grpcDeliver := delivers.NewGRPCDeliver()

			postgresService := services.NewPostgreSQLService(appName)
			defer postgresService.Close()
			elasticserachService := services.NewElasticSearchService(appName)
			defer elasticserachService.Close()

			persistenceRepo := persistence.NewRepository(postgresService)
			indexerRepo := indexer.NewRepository(elasticserachService)

			schoolService := app.NewSchoolService(&persistenceRepo, &indexerRepo)
			studentService := app.NewStudentService(&persistenceRepo, &indexerRepo)
			applicationService := app.NewApplicationService(&persistenceRepo, &indexerRepo)

			unicampus_api_admission_v1alpha1.RegisterAdmissionServiceServer(grpcDeliver.Server(), &server{
				School:      schoolService,
				Student:     studentService,
				Application: applicationService,
			})

			grpcDeliver.Run()
		},
	}

	return cmd
}
