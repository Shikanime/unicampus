package unictl

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/spf13/cobra"
	unicampus_api_education_v1alpha1 "gitlab.com/deva-hub/unicampus/api/education/v1alpha1"
)

func NewSeedSchoolCommand(client unicampus_api_education_v1alpha1.AdmissionServiceClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "school [uuid] [name] [description]",
		Short: "Manage a School ressource",
		Long:  `school command is used for managing a school type ressource.`,
		Run: func(cmd *cobra.Command, args []string) {
			schoolData, err := client.RegisterSchool(context.Background(), &unicampus_api_education_v1alpha1.School{
				UUID:        "yo",
				Name:        "ETNA",
				Description: "Desc",
				Locations: []*unicampus_api_education_v1alpha1.Location{
					&unicampus_api_education_v1alpha1.Location{
						Address: "bo",
					},
				},
			})
			if err != nil {
				fmt.Printf("education/school creation failed: %s", err)
			} else {
				fmt.Printf("education/school created: %s", schoolData)
			}
		},
	}

	return cmd
}

func NewCreateSchoolCommand(client unicampus_api_education_v1alpha1.AdmissionServiceClient) *cobra.Command {
	uuidFlag := uuid.New().String()
	descriptionFlag := "Pas de description"

	cmd := &cobra.Command{
		Use:   "school [name]",
		Short: "Manage a School ressource",
		Long:  `school command is used for managing a school type ressource.`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("name is required")
			}

			if len(args[0]) < 3 {
				return fmt.Errorf("school name require at least 3 characters but got: %s", args[0])
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(uuidFlag)
			schoolData, err := client.RegisterSchool(context.Background(), &unicampus_api_education_v1alpha1.School{
				UUID:        uuid.New().String(),
				Name:        args[0],
				Description: "Pas de description",
			})
			if err != nil {
				fmt.Printf("education/school creation failed: %s", err)
			} else {
				fmt.Printf("education/school created: %s", schoolData)
			}
		},
	}

	cmd.Flags().StringVarP(&uuidFlag, "uuid", "u", "", "Set school UUID")
	cmd.Flags().StringVarP(&descriptionFlag, "description", "d", "", "Set school description")

	return cmd
}
