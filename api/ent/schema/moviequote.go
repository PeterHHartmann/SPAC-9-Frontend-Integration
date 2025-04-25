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
		field.String("quote").
			NotEmpty().
			Unique(),
		field.String("context"),
	}
}

// Mixin for the audit fields (created_at, updated_at)
func (MovieQuote) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// Edges of the MovieQuote i.e. relationships.
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
		edge.From("language", Language.Type).
			Ref("quotes").
			Unique(),
			
	}
}
