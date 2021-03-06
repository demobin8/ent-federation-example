package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"user/ent"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user UserInput) (*ent.User, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("gin context get header accept: %s\n", gc.GetHeader("accept"))

	client := ent.FromContext(ctx)
	return client.User.Create().
		SetAge(*user.Age).
		SetName(user.Name).
		SetUsername(user.Username).Save(ctx)
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder) (*ent.UserConnection, error) {
	return r.client.User.Query().Paginate(ctx, after, first, before, last, ent.WithUserOrder(orderBy))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
