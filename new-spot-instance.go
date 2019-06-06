package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	// Create EC2 service client
	svc := ec2.New(session.New())
	input := &ec2.RequestSpotInstancesInput{
		InstanceCount: aws.Int64(1),
		LaunchSpecification: &ec2.RequestSpotLaunchSpecification{
			IamInstanceProfile: &ec2.IamInstanceProfileSpecification{
				Arn: aws.String("arn:aws:iam::033940273928:instance-profile/gosdk"),
			},
			ImageId:      aws.String("ami-0cb72367e98845d43"),
			InstanceType: aws.String("t2.micro"),
			KeyName:      aws.String("abryant-us-west-2"),
			SecurityGroupIds: []*string{
				aws.String("sg-253f426d"),
			},
			SubnetId: aws.String("subnet-e68fbf9f"),
		},
		SpotPrice: aws.String("0.050"),
		Type:      aws.String("one-time"),
	}

	result, err := svc.RequestSpotInstances(input)
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

	fmt.Println(result)
}
