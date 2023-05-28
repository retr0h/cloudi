package main

import (
	"context"

	ec2sdk "github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/retr0h/cloudi/pkg/clients/aws"
	"github.com/retr0h/cloudi/pkg/clients/gcp"
)

func main() {
	ctx := context.TODO()
	a := aws.NewClient()

	var _ aws.AWSProvider = (*aws.Client)(nil)

	// ec2
	ec2, _ := a.EC2("us-east-1")
	ec2.CreateInstance("foo")
	ec2.DeleteInstance("foo")
	ec2.GetInstanceStatus("foo")

	// ec2 vpc
	vpc, _ := a.VPC("us-east-1")
	vpc.CreateVpc(ctx, &ec2sdk.CreateVpcInput{})

	var _ gcp.GCPProvider = (*gcp.Client)(nil)

	g := gcp.NewClient()
	compute, _ := g.Compute()
	compute.CreateInstance("bar")
	compute.CreateInstance("bar")
	compute.CreateInstance("bar")
}
