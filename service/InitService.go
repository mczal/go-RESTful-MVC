package service

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// You could change this using your own DB

var svc *dynamodb.DynamoDB // mczal: this variable will be visible to all services in package service

// Make sure this func is only called once from Main.go

func GenerateDynamoDBSvc() {
	// sess := session.Must(session.NewSession())
	// svc = dynamodb.New(sess)
}
