package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// BaseNavState holds the schema definition for the BaseNavState entity.
type BaseNavState struct {
	ent.Schema
}

func (BaseNavState) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IdMixin{},
		CreatedAtMixin{},
		DeletedAtMixin{},
		UpdatedAtMixin{},
	}
}

// Fields of the BaseNavState.
func (BaseNavState) Fields() []ent.Field {
	return []ent.Field{
		field.String("base_id"),
		field.String("user_id"),
		field.Int("index"),
	}
}

// Edges of the BaseNavState.
func (BaseNavState) Edges() []ent.Edge {
	return nil
}
