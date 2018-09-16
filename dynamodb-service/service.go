package main

import (
	dyn "github.com/jensoncs/go-zero-to-hero/dynamodb-service/dynamodb"
)

func main() {
	dyn.GetZipCodeItemsFromDynamoDb("0212xyz")
}
