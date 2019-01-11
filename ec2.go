package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	sess := session.Must(session.NewSession())
	ec2svc := ec2.New(sess, aws.NewConfig().WithRegion("ap-northeast-1"))

	ec2svc.CreateVpc(&ec2.CreateVpcInput{
		CidrBlock:                   aws.String("10.0.0.0/16"),
		AmazonProvidedIpv6CidrBlock: aws.Bool(false),
		InstanceTenancy:             aws.String("default"),
	})

	fmt.Println(ec2svc.DescribeAccountAttributes(nil))
	fmt.Println(ec2svc.DescribeNatGateways(nil))
}
