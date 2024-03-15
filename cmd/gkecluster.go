package cmd

import (
	"fmt"
	"infop/internal/gke"
	"os"

	"github.com/spf13/cobra"
)

var (
	clusterName string
	projectID   string
	region      string
	nodeCount   int
	machineType string
)

var gkeClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Create a GKE cluster",
	Long:  `Create a Google Kubernetes Engine (GKE) cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		if clusterName == "" || projectID == "" || region == "" {
			fmt.Println("Cluster name, project ID, and region must be specified with -n, -p, and -r flags respectively")
			return
		}
		if nodeCount <= 0 {
			nodeCount = 1 // Default to 1 if an invalid or no count is provided
		}
		if machineType == "" {
			machineType = "e2-medium" // Default machine type if not specified
		}
		if err := gke.CreateCluster(projectID, clusterName, region, nodeCount, machineType); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating GKE cluster: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	createCmd.AddCommand(gkeClusterCmd)
	gkeClusterCmd.Flags().StringVarP(&clusterName, "name", "n", "", "Name of the GKE cluster")
	gkeClusterCmd.Flags().StringVarP(&projectID, "project", "p", "", "GCP Project ID where the cluster will be created")
	gkeClusterCmd.Flags().StringVarP(&region, "region", "r", "", "Region where the cluster will be created")
	gkeClusterCmd.Flags().IntVarP(&nodeCount, "nodes", "c", 1, "Number of nodes for the GKE cluster")
	gkeClusterCmd.Flags().StringVarP(&machineType, "machine", "m", "", "Machine type for the GKE cluster nodes (e.g., e2-medium)")
}
