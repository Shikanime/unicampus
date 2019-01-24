package commands

import (
	"github.com/spf13/cobra"
	unicampus_api_admission_v1alpha1 "gitlab.com/deva-hub/unicampus/api/admission/v1alpha1"
	app "gitlab.com/deva-hub/unicampus/internal/app/admission"
	"gitlab.com/deva-hub/unicampus/internal/app/admission/repositories/indexer"
	"gitlab.com/deva-hub/unicampus/internal/app/admission/repositories/persistence"
	"gitlab.com/deva-hub/unicampus/internal/pkg/delivers"
	"gitlab.com/deva-hub/unicampus/internal/pkg/services"
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
