package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	dynamoTableName       = "calender-service-zip-codes.dev"
	dynamoTablePrimaryKey = "lead_cus_code"
)

type ZipCodes struct {
	Group       string   `json:"Group"`
	LeadCusCode string   `json:"lead_cus_code"`
	Location    string   `json:"Location"`
	Locations   []string `json:"Locations"`
	Timezone    string   `json:"Timezone"`
	ZipCode     string   `json:"ZipCode"`
	Zone        string   `json:"zone"`
}

func main() {
	getZipCodeItemsFromDynamoDb("0212xyz")
}

func getZipCodeItemsFromDynamoDb(zipcode string) {
	dynamoRe := ZipCodes{}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1"),
	},
	)
	svc := dynamodb.New(sess)
	result1, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result1)
	params := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			dynamoTablePrimaryKey: {
				S: aws.String(zipcode),
			},
		},
		TableName:      aws.String(dynamoTableName),
		ConsistentRead: aws.Bool(true),
	}

	result, err := svc.GetItem(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	err = dynamodbattribute.UnmarshalMap(result.Item, &dynamoRe)
	if err != nil {
		fmt.Println("Error retriving data")
	}
	fmt.Println(dynamoRe)
}
