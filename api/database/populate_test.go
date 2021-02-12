package database

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"testing"

	"github.com/stretchr/testify/require"
)

func (m *mockDynamo) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return nil, nil
}

func TestPopulate(t *testing.T) {
	dynamo := &mockDynamo{}
	dal := DataLoader{
		BaseUrl:        "http://test.com",
		ImageUrl:       "http://test.com",
		DynamodbClient: dynamo,
	}
	err := dal.Populate()
	require.Equal(t, err, nil)
}
