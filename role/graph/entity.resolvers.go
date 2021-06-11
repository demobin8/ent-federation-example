package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"role/ent"
	"role/ent/userrole"
)

func (r *entityResolver) FindPermissionByID(ctx context.Context, id int) (*ent.Permission, error) {
	return r.client.Permission.Get(ctx, id)
}

func (r *entityResolver) FindRoleByID(ctx context.Context, id int) (*ent.Role, error) {
	return r.client.Role.Get(ctx, id)
}

func (r *entityResolver) FindUserByID(ctx context.Context, id int) (*User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *entityResolver) FindUserRoleByUserIDAndRoleID(ctx context.Context, userID int, roleID int) (*ent.UserRole, error) {
	return r.client.UserRole.Query().Where(userrole.UserIdEQ(userID), userrole.RoleIdEQ(roleID)).First(ctx)
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
