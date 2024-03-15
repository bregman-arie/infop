package gcp

import (
	"context"
	"fmt"
	"time"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
)

func CreateProject(projectName string) error {
	ctx := context.Background()
	client, err := resourcemanager.NewProjectsClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create client: %v", err)
	}
	defer client.Close()

	// Generate a unique project ID using the project name and current timestamp
	projectID := fmt.Sprintf("%s-%d", projectName, time.Now().Unix())

	req := &resourcemanagerpb.CreateProjectRequest{
		Project: &resourcemanagerpb.Project{
			ProjectId:   projectID,
			DisplayName: projectName,
		},
	}

	op, err := client.CreateProject(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to create project: %v", err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("failed to finalize project creation: %v", err)
	}

	fmt.Printf("Project '%s' with ID '%s' created successfully\n", projectName, projectID)
	return nil
}
