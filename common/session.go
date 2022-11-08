package common

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
)

func BuildSession() *session.Session {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("[Session] Error loading .env file")
	}

	region := os.Getenv("AWS_REGION")
	accessKey := os.Getenv("AWS_ACCESS_KEY")
	secretKey := os.Getenv("AWS_SECRET_KEY")
	sessionConfig := &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	}

	sess, err := session.NewSession(sessionConfig)
	if err != nil {
		log.Fatalf("Error to create a new session")
	}
	return sess
}
