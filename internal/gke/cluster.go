package gke

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"google.golang.org/api/compute/v1"
	"google.golang.org/api/container/v1"
)

func getRandomZone(ctx context.Context, computeService *compute.Service, projectID, region string) (string, error) {
	resp, err := computeService.Zones.List(projectID).Filter(fmt.Sprintf("region eq .*%s$", region)).Do()
	if err != nil {
		return "", fmt.Errorf("failed to list zones: %v", err)
	}

	if len(resp.Items) == 0 {
		return "", fmt.Errorf("no zones found in region %s", region)
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(resp.Items))
	return resp.Items[randomIndex].Name, nil
}

func CreateCluster(projectID, clusterName, region string, nodeCount int, machineType string) error {
	ctx := context.Background()

	computeService, err := compute.NewService(ctx)
	if err != nil {
		return fmt.Errorf("compute.NewService: %v", err)
	}

	zone, err := getRandomZone(ctx, computeService, projectID, region)
	if err != nil {
		return fmt.Errorf("getRandomZone: %v", err)
	}

	gkeService, err := container.NewService(ctx)
	if err != nil {
		return fmt.Errorf("container.NewService: %v", err)
	}

	clusterConfig := &container.CreateClusterRequest{
		Cluster: &container.Cluster{
			Name:             clusterName,
			InitialNodeCount: int64(nodeCount),
			NodeConfig: &container.NodeConfig{
				MachineType: machineType,
			},
		},
	}

	_, err = gkeService.Projects.Zones.Clusters.Create(projectID, zone, clusterConfig).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("failed to create cluster: %v", err)
	}

	fmt.Printf("Cluster %s with %d nodes (%s) creation initiated in project %s, region %s, zone %s successfully\n", clusterName, nodeCount, machineType, projectID, region, zone)
	return nil
}
