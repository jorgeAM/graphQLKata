package graphql

import "github.com/jorgeAM/graphqlKata/internal/platform/pg"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserRepo *pg.UserPostgresRepository
	TodoRepo *pg.TodoPostgresRepository
}
