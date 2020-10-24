package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "go-file-server",
	Version: "0.1.0",
	Short:   "A file server built in Go using Fiber",
}

// Execute executes the root command.
func Execute() error {
	rootCmd.AddCommand(startCommand)
	return rootCmd.Execute()
}
