package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jorgeAM/graphqlKata/internal/platform/graphql"
	"github.com/jorgeAM/graphqlKata/internal/platform/pg"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := pg.NewClient()

	if err != nil {
		panic(err)
	}

	userRepository := pg.NewUserPostgresRepository(db)
	todoRepository := pg.NewTodoPostgresRepository(db)

	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{
		Resolvers: &graphql.Resolver{
			UserRepo: userRepository,
			TodoRepo: todoRepository,
		},
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
