package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	//"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/retr0h/cloudi/pkg/clients/aws/ec2/ec2"
	"github.com/retr0h/cloudi/pkg/clients/aws/ec2/vpc"
)

// AWSProvider interface for the AWS client
type AWSProvider interface {
	// EC2 is the external client to manage EC2 resources
	EC2(region string) (ec2.Client, error)
	// VPC is the external client to manage VPC resources
	VPC(region string) (vpc.Client, error)
}

// Client has aws methods implemented
type Client struct{}

// NewClient returns a new client
func NewClient() *Client {
	return &Client{}
}

// EC2 is the external client to manage EC2 resources
func (c *Client) EC2(region string) (ec2.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, err
	}

	return ec2.NewExternal(cfg), nil
}

// VPC is the external client to manage VPC resources
func (c *Client) VPC(region string) (vpc.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, err
	}

	return vpc.NewExternal(cfg), nil
}
