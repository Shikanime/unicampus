package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	unicampus_api_education_v1alpha1 "gitlab.com/deva-hub/unicampus/api/education/v1alpha1"
	"gitlab.com/deva-hub/unicampus/cmd/unictl/app"
	"google.golang.org/grpc"
)

var (
	name = "unictl"
	host = flag.String("host", "localhost", "API host")
	port = flag.Uint("port", 50051, "API port")
	url  = flag.String("url", fmt.Sprintf("%s:%d", *host, *port), "API URL entrypoint")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Fail to connect: %s", err)
	}

	client := unicampus_api_education_v1alpha1.NewAdmissionServiceClient(conn)

	cmdCreate := &cobra.Command{
		Use:   "create [ressource name]",
		Short: "Create a ressource",
		Long:  "create command is used for any ressource creation.",
	}

	cmdSeed := &cobra.Command{
		Use:   "seed [ressource name]",
		Short: "Seed a ressource",
		Long:  "seed command is used for to bootstap with basic ressource.",
	}

	cmdSchool := app.NewCreateSchoolCommand(client)
	cmdSeedSchool := app.NewSeedSchoolCommand(client)

	cmd := &cobra.Command{Use: name}

	cmd.AddCommand(cmdCreate, cmdSeed)
	cmdCreate.AddCommand(cmdSchool)
	cmdSeed.AddCommand(cmdSeedSchool)

	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error when exiting: %s", err)
	}
}
