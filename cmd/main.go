package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"example.com/go-openstack/v2/internal/openstack"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	// Command line flags for basic server actions
	startID := flag.String("start", "", "ID of server to start")
	stopID := flag.String("stop", "", "ID of server to stop")
	rebootID := flag.String("reboot", "", "ID of server to reboot")
	snapshotID := flag.String("snapshot", "", "ID of server to snapshot")
	snapshotName := flag.String("snapshot-name", "", "Name of image to create")
	flag.Parse()

	ctx := context.Background()

	client, err := openstack.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create OpenStack client: %v", err)
	}

	if *startID != "" {
		if err := client.StartServer(ctx, *startID); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
		fmt.Printf("Started server %s\n", *startID)
	}

	if *stopID != "" {
		if err := client.StopServer(ctx, *stopID); err != nil {
			log.Fatalf("Failed to stop server: %v", err)
		}
		fmt.Printf("Stopped server %s\n", *stopID)
	}

	if *rebootID != "" {
		if err := client.RebootServer(ctx, *rebootID); err != nil {
			log.Fatalf("Failed to reboot server: %v", err)
		}
		fmt.Printf("Rebooted server %s\n", *rebootID)
	}

	if *snapshotID != "" {
		name := *snapshotName
		if name == "" {
			name = fmt.Sprintf("%s-snapshot", *snapshotID)
		}
		imageID, err := client.CreateImage(ctx, *snapshotID, name)
		if err != nil {
			log.Fatalf("Failed to create image: %v", err)
		}
		fmt.Printf("Created image %s from server %s\n", imageID, *snapshotID)
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

	// Delete the newly created server
	fmt.Printf("Deleting server %s...\n", newServer.ID)
	if err := client.DeleteServer(ctx, newServer.ID); err != nil {
		log.Fatalf("Failed to delete server: %v", err)
	}
	fmt.Println("Server deleted")
}
