package main

import (
	"context"
	"fmt"
	"log"

	"example.com/go-openstack/v2/internal/openstack"
)

func main() {
	ctx := context.Background()

	client, err := openstack.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create OpenStack client: %v", err)
	}

	// List servers example
	servers, err := client.ListServers(ctx)
	if err != nil {
		log.Fatalf("Failed to list servers: %v", err)
	}

	fmt.Println("Existing servers:")
	for i, server := range servers {
		fmt.Printf("  server %d: id=%s, name=%s\n", i, server.ID, server.Name)
	}

	// Create server example
	config := openstack.ServerConfig{
		Name:           "test-server",
		FlavorID:       "t2.micro",               // Replace with actual flavor ID
		ImageID:        "ubuntu-20.04",           // Replace with actual image ID
		NetworkIDs:     []string{"network-uuid"}, // Replace with actual network ID
		SecurityGroups: []string{"default"},
	}

	newServer, err := client.CreateServer(ctx, config)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	fmt.Printf("Created new server: id=%s, name=%s\n", newServer.ID, newServer.Name)
}
