package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type CreatedAtMixin struct {
	mixin.Schema
}

func (CreatedAtMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at"),
	}
}

type DeletedAtMixin struct {
	mixin.Schema
}

func (DeletedAtMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_at").Nillable(),
	}
}

type IdMixin struct {
	mixin.Schema
}

func (IdMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StructTag(`bson:"_id,omitempty"`),
	}
}

type UpdatedAtMixin struct {
	mixin.Schema
}

func (UpdatedAtMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("updated_at"),
	}
}
