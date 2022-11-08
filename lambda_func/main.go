package main

import (
	"context"
	"errors"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	msg := sqsEvent.Records[0].Body
	if msg == "" {
		return errors.New("The message is empty")
	} else {
		svc := sqs.New(session.New(), nil)

		sendInput := &sqs.SendMessageInput{
			MessageBody: aws.String(encode(msg)),
			QueueUrl:    aws.String(""),
		}

		_, err := svc.SendMessage(sendInput)
		if err != nil {
			return err
		}

		return nil
	}
}

func encode(msg string) string {
	replacer := strings.NewReplacer("A", "4", "a", "4", "E", "3", "e", "3", "I", "1", "i", "1", "O", "0", "o", "0", "S", "5", "s", "5", "L", "|", "l", "|")
	return replacer.Replace(msg)
}
