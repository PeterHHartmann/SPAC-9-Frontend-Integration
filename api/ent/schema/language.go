package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Language holds the schema definition for the Language entity.
type Language struct {
	ent.Schema
}

// Fields of the Language.
func (Language) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Unique(), // Ensure category names are unique
	}
}

// Mixin for the audit fields (created_at, updated_at)
func (Language) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// Edges of the Language.
func (Language) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("quotes", MovieQuote.Type),
	}
}
