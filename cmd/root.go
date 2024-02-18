package cmd

import (
    "log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
    Use: "diybase",
    Short: "A simple web backend",
    Long: "A simple web backend with multiple database tables",
}

func Execute() {
    rootCmd.AddCommand(serveCmd)
    rootCmd.AddCommand(initCmd)

    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}
