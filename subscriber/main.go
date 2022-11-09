package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/joho/godotenv"
	"github.com/linielson/aws-sns-sqs/common"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("[Subscriber] Error loading .env file")
	}
	email := os.Getenv("EMAIL_ADDRESS_SUB")
	sub := os.Getenv("AWS_SNS_TOPIC_ARN")
	subscribeTopic(email, sub)
}

func subscribeTopic(emailAddress string, topicArn string) {
	awsSession := common.BuildSession()
	svc := sns.New(awsSession)
	result, err := svc.Subscribe(&sns.SubscribeInput{
		Endpoint: aws.String(emailAddress),
		Protocol: aws.String("email"),
		TopicArn: aws.String(topicArn),
	})
	if err != nil {
		fmt.Println("Got an error subscribing to the topic: ", err)
		fmt.Println()
		return
	}

	fmt.Println(*result.SubscriptionArn)
}
