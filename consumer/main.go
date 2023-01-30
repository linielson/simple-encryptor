package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/joho/godotenv"
	"github.com/linielson/aws-sns-sqs/common"
	"github.com/linielson/aws-sns-sqs/consumer/decryptor"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	err := godotenv.Load("../.env")
	if err != nil {
		panic("[Consumer] Error loading .env file")
	}

	queueUrl := os.Getenv("AWS_SQS_ENCRYPTED_QUEUE")
	subscribe(queueUrl, sigs)
}

func subscribe(queueUrl string, cancel <-chan os.Signal) {
	awsSession := common.BuildSession()
	svc := sqs.New(awsSession, nil)

	for {
		messages := receiveMessages(svc, queueUrl)
		for _, msg := range messages {
			if msg == nil {
				continue
			}

			fmt.Println("Original: ", *msg.Body)
			fmt.Println("Decripted: ", decryptor.DecryptMessage(*msg.Body))
			go deleteMessage(svc, queueUrl, msg.ReceiptHandle)
		}

		select {
		case <-cancel:
			return
		case <-time.After(100 * time.Millisecond):
		}
	}
}

func receiveMessages(svc *sqs.SQS, queueUrl string) []*sqs.Message {
	receiveMessagesInput := &sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            aws.String(queueUrl),
		MaxNumberOfMessages: aws.Int64(10),
		WaitTimeSeconds:     aws.Int64(3),
		VisibilityTimeout:   aws.Int64(20),
	}

	receiveMessageOutput, err := svc.ReceiveMessage(receiveMessagesInput)
	if err != nil {
		fmt.Println("Receive Error: ", err)
		return nil
	}

	if receiveMessageOutput == nil || len(receiveMessageOutput.Messages) == 0 {
		return nil
	}

	return receiveMessageOutput.Messages
}

func deleteMessage(svc *sqs.SQS, queueUrl string, handle *string) {
	deleteInput := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueUrl),
		ReceiptHandle: handle,
	}
	_, err := svc.DeleteMessage(deleteInput)

	if err != nil {
		fmt.Println("Delete Error", err)
		return
	}
}
