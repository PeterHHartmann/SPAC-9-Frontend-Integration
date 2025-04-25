package schema

import (
	"fmt"
	"time"

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
		field.Int("year").
			// Validator for ensuring a reasonable year
			Validate(func(y int) error {
				minimum := 1888
				maximum := time.Now().Year()+2
				if y < minimum || y > maximum {
					return fmt.Errorf("year must be between %d and %d", minimum, maximum)
				}
				return nil
			}),
	}
}

// Mixin for the audit fields (created_at, updated_at)
func (Movie) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}

// Edges of the Movie i.e. relationships.
func (Movie) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).
			Ref("movies").
			Required(),
		edge.From("characters", Character.Type). 
			Ref("movies").
			Required(),
		edge.To("quotes", MovieQuote.Type),
	}
}
