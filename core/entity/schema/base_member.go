package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// BaseMember holds the schema definition for the BaseMember entity.
type BaseMember struct {
	ent.Schema
}

func (BaseMember) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IdMixin{},
		CreatedAtMixin{},
		DeletedAtMixin{},
		UpdatedAtMixin{},
	}
}

// Fields of the BaseMember.
func (BaseMember) Fields() []ent.Field {
	return []ent.Field{
		field.String("base_id"),
		field.String("user_id"),
	}
}

// Edges of the BaseMember.
func (BaseMember) Edges() []ent.Edge {
	return nil
}
