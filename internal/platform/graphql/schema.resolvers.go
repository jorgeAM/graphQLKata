package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jorgeAM/graphqlKata/internal/platform/uuid"
	"github.com/jorgeAM/graphqlKata/internal/user"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input LoginInput) (*user.User, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// SignUp is the resolver for the signUp field.
func (r *mutationResolver) SignUp(ctx context.Context, input SignUpInput) (*user.User, error) {
	user := &user.User{
		ID:       uuid.NewID(),
		Name:     input.Name,
		Surname:  input.Surname,
		Email:    input.Email,
		Password: input.Password,
	}

	err := r.UserRepo.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*user.User, error) {
	return r.UserRepo.FindByID(ctx, id)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
