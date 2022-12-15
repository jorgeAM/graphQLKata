package pg

import (
	"context"

	"github.com/jorgeAM/graphqlKata/internal/todo"
)

type TodoPostgresRepository struct {
	*Client
}

func NewTodoPostgresRepository(client *Client) *TodoPostgresRepository {
	return &TodoPostgresRepository{client}
}

func (repo *TodoPostgresRepository) Create(ctx context.Context, todo *todo.Todo) error {
	_, err := repo.NewInsert().Model(todo).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (repo *TodoPostgresRepository) FindAll(ctx context.Context) ([]*todo.Todo, error) {
	var todos []*todo.Todo
	err := repo.NewSelect().Model(&todos).Scan(ctx)

	if err != nil {
		return nil, err
	}

	return todos, nil
}
