package main

import (
	"log"

	"github.com/Shikanime/unicampus/internal/app/admission/commands"
	"github.com/spf13/cobra"
)

const (
	APP_NAME = "admission"
)

func main() {
	cmd := &cobra.Command{Use: "unicampus"}

	cmd.AddCommand(
		commands.NewStart(APP_NAME),
		commands.NewSetup(APP_NAME),
	)

	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error when exiting: %s", err)
	}
}
