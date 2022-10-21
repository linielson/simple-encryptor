package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func BuildSession() *session.Session {
	region := "YOUR REGION HERE"
	accessKey := "YOUR ACCESS KEY HERE"
	secretKey := "YOUR SECRET KEY HERE"
	sessionConfig := &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	}

	sess, err := session.NewSession(sessionConfig)
	if err != nil {
		panic(err)
	}
	println("SUCCESS")
	return sess
}
