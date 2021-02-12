package database

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/ivictorpd/pokedex-server/api/models"
	"github.com/ivictorpd/pokedex-server/graph/model"
	"testing"

	"github.com/stretchr/testify/require"
)

type mockDynamo struct {
	dynamodbiface.DynamoDBAPI
}

func (m *mockDynamo) ScanPages(*dynamodb.ScanInput, func(*dynamodb.ScanOutput, bool) bool) error {
	return nil
}

func (m *mockDynamo) UpdateItem(*dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	return nil, nil
}

func (m *mockDynamo) GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	dal := DAL{}
	sc := &ScanAttributes{}
	dal.addFilter(sc, "test", "test")
	return &dynamodb.GetItemOutput{
		ConsumedCapacity: nil,
		Item:             sc.ExpressionAttributeValues,
	}, nil
}

func TestSortMap(t *testing.T) {
	dynamo := &mockDynamo{}
	dal := DAL{
		DynamodbClient: dynamo,
	}
	pok := &models.Pokemon{}

	PokemonsOut := dal.sortMap(pok)

	require.NotEqual(t, PokemonsOut, nil)

}

func TestListPokemonTypes(t *testing.T) {
	dynamo := &mockDynamo{}
	dal := DAL{
		DynamodbClient: dynamo,
	}
	query := model.PokemonsQueryInput{
		Limit:  aws.Int(10),
		Offset: nil,
		Filter: &model.PokemonFilterInput{
			SearchType: nil,
			SearchName: nil,
			IsFavorite: aws.Bool(true),
		},
	}
	PokemonsOut, err := dal.ListPokemon(query)

	poc := &model.PokemonConnection{
		Limit: 0, Offset: "", Count: 0, Edges: []*model.Pokemon(nil),
	}

	require.Equal(t, PokemonsOut, poc)
	require.Equal(t, err, nil)
}

func TestUpdateFavoritePokemons(t *testing.T) {
	dynamo := &mockDynamo{}
	dal := DAL{
		DynamodbClient: dynamo,
	}
	PokemonsOut, err := dal.UpdateFavoritePokemon("10", aws.Bool(true))

	require.NotEqual(t, PokemonsOut, nil)
	require.Equal(t, err, nil)
}

func TestCreateFilter(t *testing.T) {
	dynamo := &mockDynamo{}
	dal := DAL{
		DynamodbClient: dynamo,
	}
	t.Run("add filter by isfavorite true", func(t *testing.T) {
		query := model.PokemonsQueryInput{
			Limit:  nil,
			Offset: nil,
			Filter: &model.PokemonFilterInput{
				SearchType: nil,
				SearchName: nil,
				IsFavorite: aws.Bool(true),
			},
		}
		sc := ScanAttributes{}
		sc.FilterExpression = aws.String("isfavorite = :isfavorite")
		sc.ExpressionAttributeValues = map[string]*dynamodb.AttributeValue{
			fmt.Sprint(":isfavorite"): {
				BOOL: aws.Bool(true),
			},
		}
		scanattr := dal.createFilter(query)
		require.Equal(t, *scanattr, sc)
	})
	t.Run("add filter by type pokemon", func(t *testing.T) {
		query := model.PokemonsQueryInput{
			Limit:  nil,
			Offset: nil,
			Filter: &model.PokemonFilterInput{
				SearchType: aws.String("test"),
				SearchName: nil,
				IsFavorite: nil,
			},
		}
		sc := ScanAttributes{}
		sc.FilterExpression = aws.String("contains(#types, :types)")
		sc.ExpressionAttributeValues = map[string]*dynamodb.AttributeValue{
			fmt.Sprint(":types"): {
				S: aws.String("test"),
			},
		}
		sc.ExpressionAttributeNames = map[string]*string{
			fmt.Sprint("#types"): aws.String("types"),
		}
		scanattr := dal.createFilter(query)
		require.Equal(t, *scanattr, sc)
	})
	t.Run("add filter by name pokemon", func(t *testing.T) {
		query := model.PokemonsQueryInput{
			Limit:  nil,
			Offset: nil,
			Filter: &model.PokemonFilterInput{
				SearchType: nil,
				SearchName: aws.String("test"),
				IsFavorite: nil,
			},
		}
		sc := ScanAttributes{}
		sc.FilterExpression = aws.String("contains(#name, :name)")
		sc.ExpressionAttributeValues = map[string]*dynamodb.AttributeValue{
			fmt.Sprint(":name"): {
				S: aws.String("test"),
			},
		}
		sc.ExpressionAttributeNames = map[string]*string{
			fmt.Sprint("#name"): aws.String("name"),
		}
		scanattr := dal.createFilter(query)
		require.Equal(t, *scanattr, sc)
	})
	t.Run("add filter by name pokemon and limit with offset", func(t *testing.T) {
		query := model.PokemonsQueryInput{
			Limit:  aws.Int(10),
			Offset: aws.String("010"),
			Filter: &model.PokemonFilterInput{
				SearchType: nil,
				SearchName: aws.String("test"),
				IsFavorite: nil,
			},
		}
		sc := ScanAttributes{}
		sc.FilterExpression = aws.String("contains(#name, :name)")
		sc.ExpressionAttributeValues = map[string]*dynamodb.AttributeValue{
			fmt.Sprint(":name"): {
				S: aws.String("test"),
			},
		}
		sc.ExpressionAttributeNames = map[string]*string{
			fmt.Sprint("#name"): aws.String("name"),
		}
		sc.offset = map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("010"),
			},
		}
		scanattr := dal.createFilter(query)
		require.Equal(t, *scanattr, sc)

	})
}
