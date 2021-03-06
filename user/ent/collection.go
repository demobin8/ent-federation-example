// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) *UserQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return u
}

func (u *UserQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *UserQuery {
	return u
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ul *UserLogQuery) CollectFields(ctx context.Context, satisfies ...string) *UserLogQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		ul = ul.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return ul
}

func (ul *UserLogQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *UserLogQuery {
	return ul
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (uoi *UserOaInfoQuery) CollectFields(ctx context.Context, satisfies ...string) *UserOaInfoQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		uoi = uoi.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return uoi
}

func (uoi *UserOaInfoQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *UserOaInfoQuery {
	return uoi
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (uoi *UserOpInfoQuery) CollectFields(ctx context.Context, satisfies ...string) *UserOpInfoQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		uoi = uoi.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return uoi
}

func (uoi *UserOpInfoQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *UserOpInfoQuery {
	return uoi
}
