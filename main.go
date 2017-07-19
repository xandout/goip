package main

import (
	"fmt"

	"flag"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	regionPtr := flag.String("region", "us-east-1", "Region")
	flag.Parse()
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            aws.Config{Region: aws.String(*regionPtr)},
	}))

	svc := ec2.New(sess)

	input := &ec2.DescribeAddressesInput{}

	result, err := svc.DescribeAddresses(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	addrs := result.Addresses

	for _, addr := range addrs {
		if *addr.Domain == "vpc" {
			if addr.AssociationId == nil {
				fmt.Printf("IP %s in VPC scope is unused. EIP Allocation ID: %s\n", *addr.PublicIp, *addr.AllocationId)
			}

		}

		if *addr.Domain == "standard" {
			if *addr.InstanceId == "" {
				fmt.Printf("IP %s in STANDARD scope is unused\n", *addr.PublicIp)
			}

		}

	}

}
