package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"user/ent"
)

func (r *entityResolver) FindUserByID(ctx context.Context, id int) (*ent.User, error) {
	return r.client.User.Get(ctx, id)
}

func (r *entityResolver) FindUserLogByID(ctx context.Context, id int) (*ent.UserLog, error) {
	return r.client.UserLog.Get(ctx, id)
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
