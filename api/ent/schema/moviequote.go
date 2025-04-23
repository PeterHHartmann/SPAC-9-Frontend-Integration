package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MovieQuote holds the schema definition for the MovieQuote entity.
type MovieQuote struct {
	ent.Schema
}

// Fields of the MovieQuote.
func (MovieQuote) Fields() []ent.Field {
	return []ent.Field{
		field.String("quote").NotEmpty(),
	}
}

// Edges of the MovieQuote.
func (MovieQuote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("movie", Movie.Type).
			Ref("quotes").
			Unique().
			Required(),
		edge.From("character", Character.Type).
			Ref("quotes").
			Unique().
			Required(),
	}
}
