package ec2

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	awssdk "github.com/aws/aws-sdk-go-v2/service/ec2"
)

// Client is the interface used for EC2 management
type Client interface {
	CreateInstance(instanceName string) error
	// DeleteInstance bar
	DeleteInstance(instanceName string) error
	// GetInstanceStatus baz
	GetInstanceStatus(instanceName string) (string, error)
}

// External holds EC2 management
type External struct {
	Client *awssdk.Client
}

// NewExternal returns a new client for EC2 management
func NewExternal(cfg config.Config) *External {
	// not sure why the type assert is needed here
	c := cfg.(aws.Config)
	client := awssdk.NewFromConfig(c)

	return &External{
		Client: client,
	}
}

// CreateInstance create instance
func (c *External) CreateInstance(instanceName string) error {
	fmt.Printf("AWS CreateInstance...%s\n", instanceName)
	// Implementation for creating an instance on AWS
	return nil
}

// DeleteInstance delete instance
func (c *External) DeleteInstance(instanceName string) error {
	fmt.Printf("AWS DeleteInstance...%s\n", instanceName)
	// Implementation for deleting an instance on AWS
	return nil
}

// GetInstanceStatus get instance
func (c *External) GetInstanceStatus(instanceName string) (string, error) {
	fmt.Printf("AWS GetInstanceStatus...%s\n", instanceName)
	// Implementation for getting instance status on AWS
	return "", nil
}
