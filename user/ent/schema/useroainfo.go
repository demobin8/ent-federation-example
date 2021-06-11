package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// UserOaInfo holds the schema definition for the UserOaInfo entity.
type UserOaInfo struct {
	ent.Schema
}

// Fields of the UserOaInfo.
func (UserOaInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("email"),
		field.String("email_pwd"),
	}
}

// Edges of the UserOaInfo.
func (UserOaInfo) Edges() []ent.Edge {
	return nil
}
