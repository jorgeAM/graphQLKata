package dataloader

import (
	"context"
	"fmt"

	gopher_dataloader "github.com/graph-gophers/dataloader/v7"
	"github.com/jorgeAM/graphqlKata/internal/platform/pg"
	"github.com/jorgeAM/graphqlKata/internal/user"
)

type UserLoader struct {
	*pg.UserPostgresRepository
}

func NewUserLoader(repo *pg.UserPostgresRepository) *UserLoader {
	return &UserLoader{
		repo,
	}
}

func (l *UserLoader) GetUsers(ctx context.Context, keys []string) []*gopher_dataloader.Result[*user.User] {
	var results []*gopher_dataloader.Result[*user.User]

	users, err := l.FindByIDs(ctx, keys)

	if err != nil {
		results = append(results, &gopher_dataloader.Result[*user.User]{Data: nil, Error: err})

		return results
	}

	userByID := make(map[string]*user.User)

	for _, user := range users {
		userByID[user.ID] = user
	}

	for _, key := range keys {
		u, ok := userByID[key]

		if ok {
			results = append(results, &gopher_dataloader.Result[*user.User]{Data: u, Error: nil})
		} else {
			err := fmt.Errorf("user not found %s", key)

			results = append(results, &gopher_dataloader.Result[*user.User]{Data: nil, Error: err})
		}
	}

	return results
}
