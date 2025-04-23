package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Movie holds the schema definition for the Movie entity.
type Movie struct {
	ent.Schema
}

// Fields of the Movie.
func (Movie) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().Unique(),
		field.Int("year"),
	}
}

// Edges of the Movie.
func (Movie) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("quotes", MovieQuote.Type),
	}
}
