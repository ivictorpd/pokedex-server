package database

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/ivictorpd/pokedex-server/api/models"
	"strings"
)

func getPokemonJSON() []byte {
	return []byte(models.JSON)
}

type DataLoader struct {
	BaseUrl        string
	ImageUrl       string
	DynamodbClient dynamodbiface.DynamoDBAPI
}

func (d DataLoader) Populate() error {
	var Pokemons models.Pokemon

	err := json.Unmarshal(getPokemonJSON(), &Pokemons)
	if err != nil {
		return err
	}

	for _, pokemon := range Pokemons {
		pokemon.IsFavorite = false
		pokemon.ImageUrl =
			fmt.Sprint(d.ImageUrl,
				strings.Replace(strings.ToLower(pokemon.Name), " ", "-", -1), ".jpg")

		pokemon.Sound = fmt.Sprint(d.BaseUrl, "/sounds/", pokemon.ID)
		info, err := dynamodbattribute.MarshalMap(pokemon)
		if err != nil {
			return err
		}
		input := &dynamodb.PutItemInput{
			Item:      info,
			TableName: aws.String(TableName),
		}
		_, err = d.DynamodbClient.PutItem(input)
		if err != nil {
			return err
		}
	}
	fmt.Printf("We have processed %v records\n", len(Pokemons))
	return nil
}
