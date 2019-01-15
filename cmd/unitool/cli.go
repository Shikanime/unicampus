package main

import (
	"context"
	"errors"
	"fmt"

	unicampus_api_admission_v1alpha1 "github.com/Shikanime/unicampus/api/admission/v1alpha1"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func NewCreateSchoolCommand(client unicampus_api_admission_v1alpha1.AdmissionServiceClient) *cobra.Command {
	var uuid string
	var description string

	cmd := &cobra.Command{
		Use:   "school [uuid] [name] [description]",
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
			schoolData, err := client.RegisterSchool(context.Background(), &unicampus_api_admission_v1alpha1.School{
				UUID:        uuid,
				Name:        args[0],
				Description: description,
			})
			if err != nil {
				fmt.Printf("admission/school creation failed: %s", err)
			} else {
				fmt.Printf("admission/school created: %s", schoolData)
			}
		},
	}

	cmd.Flags().StringVarP(&uuid, "uuid", "u", "", "Set school UUID")
	cmd.Flags().StringVarP(&description, "description", "d", "", "Set school description")

	return cmd
}

func main() {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	client := unicampus_api_admission_v1alpha1.NewAdmissionServiceClient(conn)

	var cmdCreate = &cobra.Command{
		Use:   "create [string to print]",
		Short: "Create a ressource",
		Long:  `create command is used for any ressource creation.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			schoolData, _ := client.RegisterSchool(context.Background(), &unicampus_api_admission_v1alpha1.School{
				UUID: "yo",
			})

			fmt.Println(schoolData)
		},
	}

	var cmdSchool = NewCreateSchoolCommand(client)

	var rootCmd = &cobra.Command{Use: "unicampus"}
	rootCmd.AddCommand(cmdCreate)
	cmdCreate.AddCommand(cmdSchool)
	rootCmd.Execute()
}
