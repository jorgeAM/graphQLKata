package user

import "context"

type Repository interface {
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id string) (*User, error)
	FindByIDs(ctx context.Context, ids []string) ([]*User, error)
}
