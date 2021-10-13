package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/dfalgout/prismagqlgen/graph/generated"
	"github.com/dfalgout/prismagqlgen/graph/model"
	"github.com/dfalgout/prismagqlgen/prisma/db"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*db.UserModel, error) {
	return r.client.User.CreateOne(
		db.User.Email.Set(input.Email),
		db.User.Password.Set(input.Password),
		db.User.FirstName.Set(input.FirstName),
		db.User.LastName.Set(input.LastName),
		db.User.Roles.Set([]string{"USER"}),
	).Exec(ctx)
}

func (r *queryResolver) Users(ctx context.Context) ([]db.UserModel, error) {
	return r.client.User.FindMany().Exec(ctx)
}

func (r *queryResolver) User(ctx context.Context, id string) (*db.UserModel, error) {
	return r.client.User.FindFirst(db.User.ID.Equals(id)).Exec(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
