package database

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/ivictorpd/pokedex-server/api/models"
	"github.com/ivictorpd/pokedex-server/graph/model"

	"sort"
)

type DAL struct {
	DynamodbClient dynamodbiface.DynamoDBAPI
}

type ScanAttributes struct {
	FilterExpression          *string
	ExpressionAttributeValues map[string]*dynamodb.AttributeValue
	ExpressionAttributeNames  map[string]*string
	offset                    map[string]*dynamodb.AttributeValue
}

func (d *DAL) sortMap(pokemons *models.Pokemon) []string {
	sTypes := make(map[string]string)
	ret := make([]string, len(sTypes))
	//select distinct values
	//map only store unique values
	for _, pokemon := range *pokemons {
		for _, ty := range pokemon.Types {
			sTypes[ty] = ""
		}
	}
	//Convert map to slice
	for k := range sTypes {
		ret = append(ret, k)
	}
	sort.Strings(ret)
	return ret
}

func (d *DAL) ListPokemonTypes() ([]string, error) {
	var ret []string
	err := d.DynamodbClient.ScanPages(&dynamodb.ScanInput{
		ProjectionExpression: aws.String("types"),
		Select:               aws.String("SPECIFIC_ATTRIBUTES"),
		TableName:            aws.String(TableName),
		TotalSegments:        nil,
	}, func(page *dynamodb.ScanOutput, last bool) bool {
		pokemons := models.Pokemon{}
		err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &pokemons)
		if err != nil {
			panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
		}
		ret = d.sortMap(&pokemons)
		return false
	})
	return ret, err
}

func (d *DAL) addFilter(attr *ScanAttributes, nameField string, valueType interface{}) *ScanAttributes {
	switch valueType.(type) {
	case bool:
		attr.FilterExpression = aws.String(fmt.Sprint(nameField, " = :", nameField))
		attr.ExpressionAttributeValues = map[string]*dynamodb.AttributeValue{
			fmt.Sprint(":", nameField): {
				BOOL: aws.Bool(valueType.(bool)),
			},
		}
	case string:
		attr.ExpressionAttributeNames = map[string]*string{
			fmt.Sprint("#", nameField): aws.String(nameField),
		}
		attr.FilterExpression = aws.String(fmt.Sprint("contains(#", nameField, ", :", nameField, ")"))
		attr.ExpressionAttributeValues = map[string]*dynamodb.AttributeValue{
			fmt.Sprint(":", nameField): {
				S: aws.String(valueType.(string)),
			},
		}
	default:
		return nil
	}
	return attr
}

func (d *DAL) createFilter(query model.PokemonsQueryInput) *ScanAttributes {
	attr := &ScanAttributes{}

	if query.Offset != nil && len(*query.Offset) > 0 {
		attr.offset = map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(*query.Offset),
			},
		}
	}
	if query.Filter == nil {
		return nil
	}

	if query.Filter.SearchType != nil {
		//Filter Pokemon by type
		d.addFilter(attr, "types", *query.Filter.SearchType)
	}
	if query.Filter.SearchName != nil {
		//Filter Pokemon by name
		d.addFilter(attr, "name", *query.Filter.SearchName)
	}
	if query.Filter.IsFavorite != nil {
		//Filter favorite
		d.addFilter(attr, "isfavorite", *query.Filter.IsFavorite)
	}

	return attr
}

func (d *DAL) ListPokemon(query model.PokemonsQueryInput) (*model.PokemonConnection, error) {

	attr := d.createFilter(query)

	recs := model.PokemonConnection{}
	err := d.DynamodbClient.ScanPages(&dynamodb.ScanInput{
		ExclusiveStartKey:         attr.offset,
		ExpressionAttributeNames:  attr.ExpressionAttributeNames,
		ExpressionAttributeValues: attr.ExpressionAttributeValues,
		FilterExpression:          attr.FilterExpression,
		Limit:                     aws.Int64(int64(*query.Limit)),
		Select:                    aws.String("ALL_ATTRIBUTES"),
		TableName:                 aws.String(TableName),
	}, func(page *dynamodb.ScanOutput, last bool) bool {
		var edge []*model.Pokemon
		err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &edge)

		for _, po := range edge {
			recs.Edges = append(recs.Edges, po)
		}
		recs.Count = int(*page.Count)

		for _, po := range page.LastEvaluatedKey {
			recs.Offset = *po.S
		}
		if query.Limit != nil {
			recs.Limit = *query.Limit
			if len(recs.Edges) == *query.Limit && recs.Edges != nil {
				return false
			}
		}
		if err != nil {
			panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
		}
		return true
	})

	return &recs, err
}

func (d *DAL) UpdateFavoritePokemon(id string, isFavorite *bool) (*model.Pokemon, error) {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				BOOL: isFavorite,
			},
		},
		TableName: aws.String(TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set isfavorite = :r"),
	}
	_, err := d.DynamodbClient.UpdateItem(input)
	if err != nil {
		return nil, err
	}
	return d.GetPokemonByID(id)
}

func (d *DAL) GetPokemonByID(id string) (*model.Pokemon, error) {
	result, err := d.DynamodbClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	po := model.Pokemon{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &po)
	if err != nil {
		return nil, err
	}
	return &po, nil
}
