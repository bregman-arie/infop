package cmd

import (
	"fmt"
	"infop/internal/gcp"
	"os"

	"github.com/spf13/cobra"
)

var projectName string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a resource",
	Long:  `Create a resource in GCP.`,
}

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Create a GCP project",
	Long:  `Create a new GCP project.`,
	Run: func(cmd *cobra.Command, args []string) {
		if projectName == "" {
			fmt.Println("Project name must be specified with -n")
			return
		}
		if err := gcp.CreateProject(projectName); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating project: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	createCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(createCmd)
	projectCmd.Flags().StringVarP(&projectName, "name", "n", "", "Name of the GCP project")
}
