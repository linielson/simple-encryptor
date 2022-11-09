package main

import (
	"bufio"
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
		panic("[Publisher] Error loading .env file")
	}

	destination := os.Getenv("AWS_SNS_TOPIC_ARN_PUB")
	publishMessages(destination)
}

func publishMessages(destination string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		if text == "\n" {
			continue
		}
		sendSNS(destination, text[:len(text)-1])
	}
}

func sendSNS(destination string, message string) {
	awsSession := common.BuildSession()
	svc := sns.New(awsSession)
	pubInput := &sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(destination),
	}
	_, err := svc.Publish(pubInput)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
