package commands

import (
	"github.com/spf13/cobra"
	unicampus_api_admission_v1alpha1 "gitlab.com/deva-hub/unicampus/api/admission/v1alpha1"
	"gitlab.com/deva-hub/unicampus/cmd/admission/app"
	"gitlab.com/deva-hub/unicampus/cmd/admission/app/repositories/indexer"
	"gitlab.com/deva-hub/unicampus/cmd/admission/app/repositories/persistence"
	"gitlab.com/deva-hub/unicampus/cmd/admission/app/repositories/recommandation"
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
			neo4jService := services.NewNeo4jService(appName)
			defer neo4jService.Close()

			persistenceRepo := persistence.NewRepository(postgresService)
			indexerRepo := indexer.NewRepository(elasticserachService)
			recommandationRepo := recommandation.NewRepository(neo4jService)

			var err error
			if err = persistenceRepo.Init(); err != nil {
				panic(err)
			}
			if err = indexerRepo.Init(); err != nil {
				panic(err)
			}

			schoolService := app.NewSchool(&persistenceRepo, &indexerRepo, &recommandationRepo)
			studentService := app.NewStudent(&persistenceRepo)
			applicationService := app.NewApplication(&persistenceRepo)

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
