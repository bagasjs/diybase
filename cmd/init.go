package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func onInit(cmd *cobra.Command, args []string) {
    log.Print("Initializing environment")
}

var initCmd = &cobra.Command {
    Use: "init",
    Short: "Init application environment",
    Run: onInit,
}
