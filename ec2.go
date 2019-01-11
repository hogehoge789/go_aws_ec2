package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	sess := session.Must(session.NewSession())
	svc := ec2.New(sess, aws.NewConfig().WithRegion("ap-northeast-1"))
	fmt.Println(svc.DescribeAccountAttributes(nil))
}
