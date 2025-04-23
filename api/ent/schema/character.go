package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Character holds the schema definition for the Character entity.
type Character struct {
	ent.Schema
}

// Fields of the Character.
func (Character) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
	}
}

// Edges of the Character.
func (Character) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("quotes", MovieQuote.Type),
	}
}
