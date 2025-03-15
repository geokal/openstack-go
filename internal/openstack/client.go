package openstack

import (
	"context"
	"os"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack"
)

// Client wraps OpenStack clients
type Client struct {
	Compute *gophercloud.ServiceClient
}

// NewClient creates a new OpenStack client with compute services
func NewClient(ctx context.Context) (*Client, error) {
	opts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	providerClient, err := openstack.AuthenticatedClient(ctx, opts)
	if err != nil {
		return nil, err
	}

	computeClient, err := openstack.NewComputeV2(providerClient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		Compute: computeClient,
	}, nil
}
