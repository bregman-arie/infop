package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "infop",
	Short: "Infop is a CLI for managing GCP projects",
	Long:  `Infop is a CLI application for creating and managing GCP projects.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Here you can define flags and configuration settings.
}
