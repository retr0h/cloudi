package compute

import (
	"fmt"
)

// Client is the interface used for Compute management
type Client interface {
	CreateInstance(instanceName string) error
	// DeleteInstance bar
	DeleteInstance(instanceName string) error
	// GetInstanceStatus baz
	GetInstanceStatus(instanceName string) (string, error)
}

// External holds GCP management
type External struct{}

// NewExternal returns a new client for Compute management
func NewExternal() *External {
	return &External{}
}

// CreateInstance create instance
func (c *External) CreateInstance(instanceName string) error {
	fmt.Printf("GCP CreateInstance...%s\n", instanceName)
	// Implementation for creating an instance on AWS
	return nil
}

// DeleteInstance delete instance
func (c *External) DeleteInstance(instanceName string) error {
	fmt.Printf("GCP DeleteInstance...%s\n", instanceName)
	// Implementation for deleting an instance on AWS
	return nil
}

// GetInstanceStatus get instance
func (c *External) GetInstanceStatus(instanceName string) (string, error) {
	fmt.Printf("GCP GetInstanceStatus...%s\n", instanceName)
	// Implementation for getting instance status on AWS
	return "", nil
}

// // CreateVpc creates an AWS VPC resource
// func (c *External) CreateVpc(
// 	ctx context.Context,
// 	input *awssdk.CreateVpcInput,
// 	opts ...func(*awssdk.Options),
// ) (*awssdk.CreateVpcOutput, error) {
// 	fmt.Printf("CreateVpc...")

// 	return &awssdk.CreateVpcOutput{}, nil
// }
