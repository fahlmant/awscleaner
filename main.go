package main

import (
	"awscleaner/pkg/aws"
	"fmt"
	_ "fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/sts"
)

const (
	accessID  = ""
	secretKey = ""
	accountid = "" // it is 12 digit account id , the sub level account
)

func main() {

	// creating a new cient with us-east-1 region by default
	client, err := clientpkg.NewClient(accessID, secretKey, "", "us-east-1")
	fmt.Println("client is %v", client)
	fmt.Println("error is %v", err)

	// trying to describe instances , it will probably fail with a root account
	result, err := client.DescribeInstances(&ec2.DescribeInstancesInput{})
	if err != nil {
		fmt.Println("you cannot describe instances yet , because error \n", err)
	} else {
		fmt.Println(result)
	}

	// creating a caller identity
	calleridentity, err := client.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	fmt.Println("calleridentity is ", calleridentity)

	fmt.Println("the test account id is ", accountid)

	assumedRole, err := client.AssumeRole(&sts.AssumeRoleInput{RoleArn: aws.String("arn:aws:iam::173028519319:role/OrganizationAccountAccessRole"), RoleSessionName: aws.String("agautam")})
	fmt.Println(assumedRole)

	assumedAccessKey := *assumedRole.Credentials.AccessKeyId
	assumedSecretKey := *assumedRole.Credentials.SecretAccessKey
	assumedSessionToken := *assumedRole.Credentials.SessionToken

	// for debugging purpose only.
	fmt.Println("new access id : ", assumedAccessKey)
	fmt.Println("new secret key is", assumedSecretKey)
	fmt.Println("new session token\n\n", assumedSessionToken)

	client2, err := clientpkg.NewClient(assumedAccessKey, assumedSecretKey, assumedSessionToken, "us-east-1")
	fmt.Println("new_client is %v", client2)
	fmt.Println("error is %v", err)

	// creating a caller identity
	calleridentity2, err := client2.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	fmt.Println("new calleridentity is ", calleridentity2)

	// trying to describe instances , it will probably fail with a root account
	result2, err := client2.DescribeInstances(&ec2.DescribeInstancesInput{})
	fmt.Println("error is ", err)
	fmt.Println("EC2 instances", result2)

	// something wrong here. no error but no EC2 instances printed

}
