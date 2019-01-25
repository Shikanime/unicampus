package commands

import (
	"encoding/csv"
	"io"
	"os"

	"gitlab.com/deva-hub/unicampus/pkg/admission"

	"github.com/spf13/cobra"
	"gitlab.com/deva-hub/unicampus/cmd/admission/app/repositories/indexer"
	"gitlab.com/deva-hub/unicampus/cmd/admission/app/repositories/persistence"
	"gitlab.com/deva-hub/unicampus/internal/pkg/services"
)

func NewSetup(appName string) *cobra.Command {
	var filename string

	cmd := &cobra.Command{
		Use:   "setup [service]",
		Short: "Initialize a service",
		Long:  `setup command is used for managing sub service intialization.`,
	}

	persistenctCmd := &cobra.Command{
		Use:   "persistence",
		Short: "Initialize persistence service",
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

			data, err := os.Open(filename)
			if err != nil {
				panic(err)
			}
			defer data.Close()

			parser := csv.NewReader(data)
			for {
				record, err := parser.Read()
				if err == io.EOF {
					break
				} else if err != nil {
					panic(err)
				}
				if err := persistenceRepo.CreateSchool(&admission.School{
					Identification: admission.Identification{
						UUID: record[0],
					},
					Name:        record[1],
					Description: record[2],
				}); err != nil {
					panic(err)
				}
			}
		},
	}

	cmd.MarkFlagRequired("filename")
	cmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "filename to read from")
	cmd.AddCommand(persistenctCmd)

	return cmd
}
