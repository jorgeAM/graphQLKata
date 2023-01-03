package dataloader

import (
	"github.com/graph-gophers/dataloader/v7"
	"github.com/jorgeAM/graphqlKata/internal/platform/pg"
	"github.com/jorgeAM/graphqlKata/internal/user"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

type Loaders struct {
	UserLoader *dataloader.Loader[string, *user.User]
}

func NewLoader(userRepo *pg.UserPostgresRepository) *Loaders {
	userLoader := NewUserLoader(userRepo)

	return &Loaders{
		UserLoader: dataloader.NewBatchedLoader(userLoader.GetUsers),
	}
}
