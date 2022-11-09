package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/joho/godotenv"
	"github.com/linielson/aws-sns-sqs/common"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("[Producer] Error loading .env file")
	}

	// messages sent direct to Normal Queue are not being triggered by lambda func (only messages from SNS -> SQS Normal Queue -> Lambda func)
	// queueURL := os.Getenv("AWS_SQS_NORMAL_QUEUE")
	queueURL := os.Getenv("AWS_SQS_ENCRYPTED_QUEUE")
	sendMessage(queueURL)
}

func sendMessage(queueURL string) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		if text == "\n" {
			continue
		}
		sendSQS(queueURL, text[:len(text)-1])
	}
}

func sendSQS(queueURL string, message string) {
	awsSession := common.BuildSession()
	svc := sqs.New(awsSession, nil)
	_, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(message),
		QueueUrl:     aws.String(queueURL),
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
