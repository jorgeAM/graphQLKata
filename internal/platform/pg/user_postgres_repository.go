package pg

import (
	"context"

	"github.com/jorgeAM/graphqlKata/internal/user"
	"github.com/uptrace/bun"
)

type UserPostgresRepository struct {
	*Client
}

func NewUserPostgresRepository(client *Client) *UserPostgresRepository {
	return &UserPostgresRepository{client}
}

func (repo *UserPostgresRepository) Create(ctx context.Context, user *user.User) error {
	_, err := repo.NewInsert().Model(user).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (repo *UserPostgresRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	user := new(user.User)
	err := repo.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserPostgresRepository) FindByIDs(ctx context.Context, ids []string) ([]*user.User, error) {
	var users []*user.User

	err := repo.NewSelect().Model(&users).Where("id IN (?)", bun.In(ids)).Scan(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}
