package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ivictorpd/pokedex-server/api/database"
	"github.com/ivictorpd/pokedex-server/graph/generated"
	"github.com/ivictorpd/pokedex-server/graph/model"
)

func (r *mutationResolver) UpdateFavoritePokemon(ctx context.Context, id string, isFavorite *bool) (*model.Pokemon, error) {
	dal := database.DAL{
		DynamodbClient: database.InitSession(),
	}
	return dal.UpdateFavoritePokemon(id, isFavorite)
}

func (r *queryResolver) ListPokemon(ctx context.Context, input model.PokemonsQueryInput) (*model.PokemonConnection, error) {
	dal := database.DAL{
		DynamodbClient: database.InitSession(),
	}
	return dal.ListPokemon(input)
}

func (r *queryResolver) GetPokemonByID(ctx context.Context, id string) (*model.Pokemon, error) {
	dal := database.DAL{
		DynamodbClient: database.InitSession(),
	}
	return dal.GetPokemonByID(id)
}

func (r *queryResolver) ListPokemonTypes(ctx context.Context) ([]string, error) {
	dal := database.DAL{
		DynamodbClient: database.InitSession(),
	}
	return dal.ListPokemonTypes()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
