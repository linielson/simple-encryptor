package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/joho/godotenv"
	"github.com/linielson/aws-sns-sqs/common"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("[Publisher] Error loading .env file")
	}

	destination := os.Getenv("AWS_SNS_TOPIC_ARN_PUB")
	publishMessages(SendSNS, destination)
}

func SendSNS(session *session.Session, destination string, message string) {
	svc := sns.New(session)
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

func publishMessages(sender func(session *session.Session, destination string, message string), destination string) {
	awsSession := common.BuildSession()
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		if text == "\n" {
			continue
		}
		sender(awsSession, destination, text[:len(text)-1])
	}
}
