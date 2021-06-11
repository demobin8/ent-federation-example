package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// UserOpInfo holds the schema definition for the UserOpInfo entity.
type UserOpInfo struct {
	ent.Schema
}

// Fields of the UserOpInfo.
func (UserOpInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("phone"),
	}
}

// Edges of the UserOpInfo.
func (UserOpInfo) Edges() []ent.Edge {
	return nil
}
