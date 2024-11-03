package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
)

func main() {
	ctx := context.Background()

	// load enviromental variables for openstack , they start with "OS_"
	opts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		panic(err)
	}

	providerClient, err := openstack.AuthenticatedClient(ctx, opts)
	if err != nil {
		panic(err)
	}

	computeClient, err := openstack.NewComputeV2(providerClient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if err != nil {
		panic(err)
	}

	// use the computeClient
	listOpts := servers.ListOpts{
		AllTenants: false,
	}

	// use the computeClient
	pager, err := servers.List(computeClient, listOpts).AllPages(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}

	servers, err := servers.ExtractServers(pager)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("servers:")
	for i, server := range servers {
		fmt.Printf("  server %d: id=%s\n", i, server.ID)
	}

}
