package dynamo

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func unmarshalMap(
	fields map[string]*dynamodb.AttributeValue,
	obj interface{},
) (
	err error,
) {
	err = dynamodbattribute.UnmarshalMap(fields, &obj)
	return
}
