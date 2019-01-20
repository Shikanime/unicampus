package commands

import (
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/indexer"
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/persistence"
	"github.com/Shikanime/unicampus/internal/pkg/services"
	"github.com/spf13/cobra"
)

func NewSetup(appName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setup",
		Short: "Setup a service",
		Long:  `setup command is used for managing sub service intialization.`,
		Run: func(cmd *cobra.Command, args []string) {
			postgresService := services.NewPostgreSQLService(appName)
			defer postgresService.Close()
			elasticserachService := services.NewElasticSearchService(appName)
			defer elasticserachService.Close()

			persistenceRepo := persistence.NewRepository(postgresService)
			indexerRepo := indexer.NewRepository(elasticserachService)

			var err error
			if err = persistenceRepo.Init(); err != nil {
				panic(err)
			}
			if err = indexerRepo.Init(); err != nil {
				panic(err)
			}
		},
	}

	return cmd
}
