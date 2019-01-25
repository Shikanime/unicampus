package main

import (
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/deva-hub/unicampus/cmd/admission/app/commands"
)

var (
	name = "admission"
)

func main() {
	cmd := &cobra.Command{Use: name}

	cmd.AddCommand(
		commands.NewStart(name),
		commands.NewSetup(name),
	)

	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error when exiting: %s", err)
	}
}
