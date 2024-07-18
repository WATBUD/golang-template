package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Board holds the schema definition for the Board entity.
type Board struct {
	ent.Schema
}

func (Board) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IdMixin{},
		CreatedAtMixin{},
		DeletedAtMixin{},
		UpdatedAtMixin{},
	}
}

// Fields of the Board.
func (Board) Fields() []ent.Field {
	return []ent.Field{
		field.String("bbbb"),
		field.String("aaa"),
		field.Int("aaa"),
	}
}

// Edges of the Board.
func (Board) Edges() []ent.Edge {
	return nil
}
