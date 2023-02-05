package main

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/events/test"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalMessage(t *testing.T) {
	inputJSON := test.ReadJSONFromFile(t, "./testdata/sqs-event.json")
	var inputEvent events.SQSEvent
	if err := json.Unmarshal(inputJSON, &inputEvent); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}

	message, err := UnmarshalMessage(inputEvent)
	assert.Nil(t, err)
	assert.Equal(t, "Message from SNS", message)
}
