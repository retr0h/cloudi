package gcp

import (
	"github.com/retr0h/cloudi/pkg/clients/gcp/compute"
)

// GCPProvider interface for the GCP client
type GCPProvider interface {
	// Compute is the external client to manage Compute resources
	Compute() (compute.Client, error)
}

// Client has aws methods implemented
type Client struct{}

// NewClient returns a new client
func NewClient() *Client {
	return &Client{}
}

// Compute is the external client to manage Compute resources
func (c *Client) Compute() (compute.Client, error) {
	return compute.NewExternal(), nil
}
