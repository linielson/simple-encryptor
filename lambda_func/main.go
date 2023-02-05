package main

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Body struct {
	Signature         string                 `json:"Signature"`
	MessageID         string                 `json:"MessageId"`
	Type              string                 `json:"Type"`
	TopicArn          string                 `json:"TopicArn"` //nolint: stylecheck
	MessageAttributes map[string]interface{} `json:"MessageAttributes"`
	SignatureVersion  string                 `json:"SignatureVersion"`
	Timestamp         time.Time              `json:"Timestamp"`
	SigningCertURL    string                 `json:"SigningCertUrl"`
	Message           string                 `json:"Message"`
	UnsubscribeURL    string                 `json:"UnsubscribeUrl"`
	Subject           string                 `json:"Subject"`
}

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	message, err := UnmarshalMessage(sqsEvent)
	if err != nil {
		return errors.New("the message is empty")
	} else {
		sess, _ := session.NewSession()
		svc := sqs.New(sess, nil)

		sendInput := &sqs.SendMessageInput{
			MessageBody: aws.String(encode(message)),
			QueueUrl:    aws.String(os.Getenv("AWS_SQS_ENCRYPTED_QUEUE")),
		}

		_, err := svc.SendMessage(sendInput)
		if err != nil {
			return err
		}

		return nil
	}
}

func UnmarshalMessage(sqsEvent events.SQSEvent) (string, error) {
	jsonBody := []byte(sqsEvent.Records[0].Body)
	var body Body
	err := json.Unmarshal(jsonBody, &body)
	if err != nil {
		return "", err
	} else {
		return body.Message, nil
	}
}

func encode(msg string) string {
	replacer := strings.NewReplacer(
		"a", "/4", "e", "/3", "i", "/1", "o", "/0", "b", "/8", "c", "/6", "g", "/9", "s", "/5", "t", "/7", "z", "/2",
	)
	return replacer.Replace(msg)
}
