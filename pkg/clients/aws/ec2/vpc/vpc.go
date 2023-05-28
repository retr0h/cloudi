package vpc

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// Client is the interface used for VPC management
type Client interface {
	CreateVpc(ctx context.Context,
		input *ec2.CreateVpcInput,
		opts ...func(*ec2.Options),
	) (*ec2.CreateVpcOutput, error)
}

// External holds VPC management
type External struct {
	Client ec2.Client
}

// NewExternal returns a new client for VPC management
func NewExternal(cfg config.Config) *External {
	// not sure why the type assert is needed here
	c := cfg.(aws.Config)
	client := ec2.NewFromConfig(c)

	return &External{
		Client: *client,
	}
}

// CreateVpc creates an AWS VPC resource
func (c *External) CreateVpc(
	ctx context.Context,
	input *ec2.CreateVpcInput,
	opts ...func(*ec2.Options),
) (*ec2.CreateVpcOutput, error) {
	fmt.Printf("AWS CreateVpc...\n")

	return &ec2.CreateVpcOutput{}, nil
}
