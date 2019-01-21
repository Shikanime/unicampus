package commands

import (
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/indexer"
	"github.com/Shikanime/unicampus/internal/app/admission/repositories/persistence"
	"github.com/Shikanime/unicampus/internal/pkg/services"
	"github.com/Shikanime/unicampus/pkg/admission"
	uuid "github.com/satori/go.uuid"
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

			persistenceRepo.CreateSchool(&admission.School{
				UUID:        uuid.NewV4().String(),
				Name:        "ETNA",
				Description: "The École des technologies numériques appliquées (ETNA) is a French private school in computer science localized at Ivry-sur-Seine. Created in 2005 by Patrice Dumoucel, the school since 2006 is part of IONIS Education Group. The certification delivered by the school are recognized by the French state.",
			})

			persistenceRepo.CreateSchool(&admission.School{
				UUID:        uuid.NewV4().String(),
				Name:        "ESGI",
				Description: ".",
			})
		},
	}

	return cmd
}
