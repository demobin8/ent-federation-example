// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pe *PermissionQuery) CollectFields(ctx context.Context, satisfies ...string) *PermissionQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		pe = pe.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return pe
}

func (pe *PermissionQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *PermissionQuery {
	return pe
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (r *RoleQuery) CollectFields(ctx context.Context, satisfies ...string) *RoleQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		r = r.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return r
}

func (r *RoleQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *RoleQuery {
	return r
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ur *UserRoleQuery) CollectFields(ctx context.Context, satisfies ...string) *UserRoleQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		ur = ur.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return ur
}

func (ur *UserRoleQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *UserRoleQuery {
	return ur
}
