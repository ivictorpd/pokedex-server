package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ivictorpd/pokedex-server/api/database"
	"github.com/ivictorpd/pokedex-server/graph"
	"github.com/ivictorpd/pokedex-server/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	database.TableName = os.Getenv("TABLE_NAME")
	database.Endpoint = os.Getenv("ENTRY_POINT")
	dt := database.ClientDatabase{
		DynamodbClient: database.InitSession(),
	}
	err := dt.InitDb()
	if err == nil {
		dt := database.DataLoader{
			BaseUrl:        fmt.Sprint(os.Getenv("BASE_URL"), "/sounds/"),
			ImageUrl:       "https://img.pokemondb.net/artwork/",
			DynamodbClient: database.InitSession(),
		}
		err = dt.Populate()
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println(err.Error())
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
