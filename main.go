package main

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

func main() {

	opts := gophercloud.AuthOptions{
		IdentityEndpoint: "https://auth.cloud.ovh.net",
		Username:         "",
		Password:         "",
		DomainName:       "ovhcloud-emea",
		DomainID:         "default",
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		panic(err)
	}
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: "UK1",
	})
	if err != nil {
		panic(err)
	}

	// use the computeClient
	pager, err := servers.List(computeClient, nil).AllPages()
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
