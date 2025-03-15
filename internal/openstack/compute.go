package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/flavors"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
)

// ServerConfig holds the configuration for creating a new server
type ServerConfig struct {
	Name           string
	FlavorID       string
	ImageID        string
	NetworkIDs     []string
	SecurityGroups []string
}

// ListServers returns all servers in the current project
func (c *Client) ListServers(ctx context.Context) ([]servers.Server, error) {
	listOpts := servers.ListOpts{
		AllTenants: false,
	}

	pager, err := servers.List(c.Compute, listOpts).AllPages(ctx)
	if err != nil {
		return nil, err
	}

	return servers.ExtractServers(pager)
}

// CreateServer creates a new virtual machine
func (c *Client) CreateServer(ctx context.Context, config ServerConfig) (*servers.Server, error) {
	networks := make([]servers.Network, len(config.NetworkIDs))
	for i, netID := range config.NetworkIDs {
		networks[i] = servers.Network{
			UUID: netID,
		}
	}

	createOpts := servers.CreateOpts{
		Name:           config.Name,
		FlavorRef:      config.FlavorID,
		ImageRef:       config.ImageID,
		Networks:       networks,
		SecurityGroups: config.SecurityGroups,
	}

	server, err := servers.Create(ctx, c.Compute, createOpts, nil).Extract()
	if err != nil {
		return nil, err
	}

	return server, nil
}

// GetFlavor retrieves flavor details by ID
func (c *Client) GetFlavor(ctx context.Context, flavorID string) (*flavors.Flavor, error) {
	return flavors.Get(ctx, c.Compute, flavorID).Extract()
}
