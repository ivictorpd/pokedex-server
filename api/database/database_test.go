package database

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"testing"

	"github.com/stretchr/testify/require"
)

func (m *mockDynamo) CreateTable(*dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	return nil, nil
}

func TestCreateDatabase(t *testing.T) {
	dynamo := &mockDynamo{}
	dal := ClientDatabase{
		DynamodbClient: dynamo,
	}
	err := dal.createTable()
	require.Equal(t, err, nil)
}
