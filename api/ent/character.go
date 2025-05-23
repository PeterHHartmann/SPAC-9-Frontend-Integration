// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/character"
	"api/ent/movie"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Character is the model entity for the Character schema.
type Character struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Actor holds the value of the "actor" field.
	Actor string `json:"actor,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CharacterQuery when eager-loading is set.
	Edges            CharacterEdges `json:"edges"`
	movie_characters *int
	selectValues     sql.SelectValues
}

// CharacterEdges holds the relations/edges for other nodes in the graph.
type CharacterEdges struct {
	// Movie holds the value of the movie edge.
	Movie *Movie `json:"movie,omitempty"`
	// Quotes holds the value of the quotes edge.
	Quotes []*MovieQuote `json:"quotes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MovieOrErr returns the Movie value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CharacterEdges) MovieOrErr() (*Movie, error) {
	if e.Movie != nil {
		return e.Movie, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: movie.Label}
	}
	return nil, &NotLoadedError{edge: "movie"}
}

// QuotesOrErr returns the Quotes value or an error if the edge
// was not loaded in eager-loading.
func (e CharacterEdges) QuotesOrErr() ([]*MovieQuote, error) {
	if e.loadedTypes[1] {
		return e.Quotes, nil
	}
	return nil, &NotLoadedError{edge: "quotes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Character) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case character.FieldID:
			values[i] = new(sql.NullInt64)
		case character.FieldName, character.FieldActor:
			values[i] = new(sql.NullString)
		case character.FieldCreatedAt, character.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case character.ForeignKeys[0]: // movie_characters
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Character fields.
func (c *Character) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case character.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case character.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case character.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case character.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case character.FieldActor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field actor", values[i])
			} else if value.Valid {
				c.Actor = value.String
			}
		case character.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field movie_characters", value)
			} else if value.Valid {
				c.movie_characters = new(int)
				*c.movie_characters = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Character.
// This includes values selected through modifiers, order, etc.
func (c *Character) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryMovie queries the "movie" edge of the Character entity.
func (c *Character) QueryMovie() *MovieQuery {
	return NewCharacterClient(c.config).QueryMovie(c)
}

// QueryQuotes queries the "quotes" edge of the Character entity.
func (c *Character) QueryQuotes() *MovieQuoteQuery {
	return NewCharacterClient(c.config).QueryQuotes(c)
}

// Update returns a builder for updating this Character.
// Note that you need to call Character.Unwrap() before calling this method if this Character
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Character) Update() *CharacterUpdateOne {
	return NewCharacterClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Character entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Character) Unwrap() *Character {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Character is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Character) String() string {
	var builder strings.Builder
	builder.WriteString("Character(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("actor=")
	builder.WriteString(c.Actor)
	builder.WriteByte(')')
	return builder.String()
}

// Characters is a parsable slice of Character.
type Characters []*Character
