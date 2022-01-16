package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

//Player holds response
type Player struct {
	ID             int      `json:"int"`
	Name           string   `json:"name"`
	HighScore      int      `json:"highScore"`
	IsOnline       bool     `json:"isOnline"`
	Location       string   `json:"location"`
	LevelsUnlocked []string `json:"levelsUnlocked"`
}

var players = []Player{
	{ID: 123, Name: "Pablo", HighScore: 1100, IsOnline: true, Location: "Italy"},
	{ID: 230, Name: "Dora", HighScore: 2100, IsOnline: false, Location: "Germany"},
}

var playerObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Player",
		Fields: graphql.Fields{
			"id":             &graphql.Field{Type: graphql.Int},
			"name":           &graphql.Field{Type: graphql.String},
			"highScore":      &graphql.Field{Type: graphql.String},
			"isOnline":       &graphql.Field{Type: graphql.Boolean},
			"location":       &graphql.Field{Type: graphql.String},
			"levelsUnlocked": &graphql.Field{Type: graphql.NewList(graphql.String)},
		},
	},
)

func main() {
	//Schema
	fields := graphql.Fields{
		"players": &graphql.Field{
			Type:        graphql.NewList(playerObject),
			Description: "All Players",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return players, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{
		Name:   "Root Query",
		Fields: fields,
	}
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}
	schema, _ := graphql.NewSchema(schemaConfig)

	//Handler
	p := handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	}
	h := handler.New(&p)
	http.Handle("/graphql", h)
	http.ListenAndServe(":8000", nil)
	log.Printf("Server listening @%d GraphQL browser on path %s", 8000, "/graphql")
}
