package todo

import "context"

type Repository interface {
	Create(ctx context.Context, todo *Todo) error
	FindAll(ctx context.Context) ([]*Todo, error)
}
