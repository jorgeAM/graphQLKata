package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jorgeAM/graphqlKata/internal/platform/uuid"
	"github.com/jorgeAM/graphqlKata/internal/todo"
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

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input CreateTodoInput) (*todo.Todo, error) {
	todo := &todo.Todo{
		ID:     uuid.NewID(),
		Text:   input.Text,
		Done:   true,
		UserID: input.UserID,
	}

	err := r.TodoRepo.Create(ctx, todo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*user.User, error) {
	return r.UserRepo.FindByID(ctx, id)
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*todo.Todo, error) {
	return r.TodoRepo.FindAll(ctx)
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *todo.Todo) (*user.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
