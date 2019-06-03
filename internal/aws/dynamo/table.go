package dynamo

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	internal "github.com/jackmcguire1/go-twitch-analytics/internal/aws"
)

var svc = dynamodb.New(internal.Session)

var (
	// ErrNotFound error for missing dynamo item
	ErrNotFound = errors.New("Item Not Found")
)

// Table dynamoDB service
type Table struct {
	Name string
}

// Client init dynamoDB service
func Client(name string) *Table {
	return &Table{Name: name}
}

// Put inserts item into dynamoDB table
func (t *Table) Put(ctx context.Context, item interface{}) (result *dynamodb.PutItemOutput, err error) {
	it, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return
	}

	result, err = svc.PutItemWithContext(ctx,
		&dynamodb.PutItemInput{
			Item:      it,
			TableName: aws.String(t.Name),
		},
	)

	return
}

// Get retrieves item from dynamoDB table
func (t *Table) Get(ctx context.Context, id string, data interface{}) (err error) {
	result, err := svc.GetItemWithContext(ctx,
		&dynamodb.GetItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"id": &dynamodb.AttributeValue{
					S: aws.String(id),
				}},
			TableName: aws.String(t.Name),
		},
	)
	if err != nil {
		return
	}

	if len(result.Item) == 0 {
		err = ErrNotFound
		return
	}
	err = unmarshalMap(result.Item, &data)

	return
}
