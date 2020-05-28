package main

import (
	"awscleaner/pkg/aws"
	"fmt"
	_ "fmt"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/sts"
)

const (
	accessID  = ""
	secretKey = ""
)

func main() {

	client, err := clientpkg.NewClient(accessID, secretKey, "", "us-east-1")
	fmt.Println("client is %v", client)
	fmt.Println("error is %v", err)

	// trying to describe instances , it will probably fail with a root account
	result, err := client.DescribeInstances(&ec2.DescribeInstancesInput{})
	if err != nil {
		fmt.Println("you cannot descibe instnces yet , becasue error \n", err)
	} else {
		fmt.Println(result)
	}

	calleridentity, err := client.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	fmt.Println("calleridentity is ", calleridentity)

}
