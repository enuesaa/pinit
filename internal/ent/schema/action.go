package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Action struct {
	ent.Schema
}

func (Action) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").
			Unique().
			Immutable(),

		field.String("name").
			MaxLen(255),

		field.String("template"),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Action) Edges() []ent.Edge {
	return []ent.Edge{}
}
