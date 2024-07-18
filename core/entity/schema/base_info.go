package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// BaseInfo holds the schema definition for the BaseInfo entity.
type BaseInfo struct {
	ent.Schema
}

func (BaseInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IdMixin{},
		CreatedAtMixin{},
		DeletedAtMixin{},
		UpdatedAtMixin{},
	}
}

// Fields of the BaseInfo.
func (BaseInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("base_id"),
		field.String("name"),
		field.String("logo"),
		field.String("color"),
	}
}

// Edges of the BaseInfo.
func (BaseInfo) Edges() []ent.Edge {
	return nil
}
