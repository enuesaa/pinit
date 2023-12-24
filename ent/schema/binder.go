package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Binder struct {
	ent.Schema
}

func (Binder) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").
			Unique().
			Immutable(),
		field.String("name").
			MaxLen(255).
			Unique(),
		field.String("category").
			MaxLen(255),
		field.Time("archived_at").
			Nillable().
			Optional(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Binder) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("notes", Note.Type),
	}
}
