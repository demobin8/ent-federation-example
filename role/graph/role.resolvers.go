package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"role/ent"
	"role/ent/role"
	"role/ent/userrole"

	"github.com/jinzhu/copier"
)

func (r *mutationResolver) CreateRole(ctx context.Context, role RoleInput) (*ent.Role, error) {
	return ent.FromContext(ctx).Role.Create().
		SetName(role.Name).
		SetRemark(role.Remark).Save(ctx)
}

func (r *mutationResolver) CreatePermission(ctx context.Context, permission PermissionInput) (*ent.Permission, error) {
	return ent.FromContext(ctx).Permission.Create().
		SetName(permission.Name).
		SetCode(permission.Code).
		SetRemark(permission.Remark).Save(ctx)
}

func (r *mutationResolver) CreateUserRole(ctx context.Context, userRole UserRoleInput) (int, error) {
	rst := ent.FromContext(ctx).UserRole.Create().
		SetRoleId(userRole.RoleID).
		SetUserId(userRole.UserID).SaveX(ctx)
	return rst.ID, nil
}

func (r *mutationResolver) CreateRolePermission(ctx context.Context, rolePermission RolePermissionInput) (*int, error) {
	r.client.Role.UpdateOneID(rolePermission.RoleID).AddPermissionIDs(rolePermission.PermissionID).SaveX(ctx)
	return nil, nil
}

func (r *queryResolver) Roles(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.RoleOrder) (*ent.RoleConnection, error) {
	return r.client.Role.Query().Paginate(ctx, after, first, before, last, ent.WithRoleOrder(orderBy))
}

func (r *queryResolver) Permissions(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.PermissionOrder) (*ent.PermissionConnection, error) {
	return r.client.Permission.Query().Paginate(ctx, after, first, before, last, ent.WithPermissionOrder(orderBy))
}

func (r *queryResolver) UserRoleDetail(ctx context.Context, userID *int) (*User, error) {
	userRoleDetail := User{
		ID: *userID,
	}

	userRoles := r.client.UserRole.Query().Where(userrole.UserIdEQ(*userID)).AllX(ctx)

	var roles []*RoleDetail

	for _, userRole := range userRoles {
		roleEntity := r.client.Role.Query().Where(role.IDEQ(userRole.RoleId)).FirstX(ctx)
		rolePermissions := r.client.Role.Query().Where(role.IDEQ(userRole.RoleId)).QueryPermissions().AllX(ctx)

		roleDetail := RoleDetail{
			&ent.Role{},
			nil,
		}
		_ = copier.CopyWithOption(roleDetail.Role, roleEntity, copier.Option{IgnoreEmpty: true, DeepCopy: false})
		roleDetail.Permissions = rolePermissions

		roles = append(roles, &roleDetail)
	}

	userRoleDetail.Roles = roles

	return &userRoleDetail, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
