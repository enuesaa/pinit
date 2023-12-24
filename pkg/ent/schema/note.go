package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Note struct {
	ent.Schema
}

func (Note) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").
			Unique().
			Immutable(),

		field.Uint("binder_id").Optional(),

		field.String("publisher").
			MaxLen(255),

		field.String("comment").
			Optional(),

		field.String("content").
			Optional(),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Note) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("binder", Binder.Type).
			Ref("notes").
			Field("binder_id").
			Unique(),
	}
}
