package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Unique(), // Ensure category names are unique
	}
}

// Mixin for the audit fields (created_at, updated_at)
func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// Edges of the Category i.e. relationships.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("movies", Movie.Type),
	}
}
