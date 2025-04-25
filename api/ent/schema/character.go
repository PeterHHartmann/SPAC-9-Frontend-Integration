package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Character holds the schema definition for the Character entity.
type Character struct {
	ent.Schema
}

// Fields of the Character.
func (Character) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("actor").
			NotEmpty(),
	}
}

// Mixin for the audit fields (created_at, updated_at)
func (Character) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// Edges of the Character i.e. relationships.
func (Character) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("movies", Movie.Type),
	}
}


// Indexes to define compound primary key
func (Character) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "actor").
			Unique(),
	}
}