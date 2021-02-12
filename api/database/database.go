package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func (c *ClientDatabase) InitDb() error {
	return c.createTable()
}

var TableName string
var Endpoint string

type ClientDatabase struct {
	DynamodbClient dynamodbiface.DynamoDBAPI
}

func InitSession() *dynamodb.DynamoDB {
	config := &aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String(Endpoint),
	}
	sess := session.Must(session.NewSession(config))
	return dynamodb.New(sess)
}

func (c *ClientDatabase) createTable() error {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(TableName),
	}

	_, err := c.DynamodbClient.CreateTable(input)
	if err != nil {
		return err
	}
	return nil
}
