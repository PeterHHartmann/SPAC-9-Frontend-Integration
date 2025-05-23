// Code generated by ent, DO NOT EDIT.

package movie

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the movie type in the database.
	Label = "movie"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldYear holds the string denoting the year field in the database.
	FieldYear = "year"
	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "category"
	// EdgeCharacters holds the string denoting the characters edge name in mutations.
	EdgeCharacters = "characters"
	// EdgeQuotes holds the string denoting the quotes edge name in mutations.
	EdgeQuotes = "quotes"
	// Table holds the table name of the movie in the database.
	Table = "movies"
	// CategoryTable is the table that holds the category relation/edge.
	CategoryTable = "movies"
	// CategoryInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoryInverseTable = "categories"
	// CategoryColumn is the table column denoting the category relation/edge.
	CategoryColumn = "category_movies"
	// CharactersTable is the table that holds the characters relation/edge.
	CharactersTable = "characters"
	// CharactersInverseTable is the table name for the Character entity.
	// It exists in this package in order to avoid circular dependency with the "character" package.
	CharactersInverseTable = "characters"
	// CharactersColumn is the table column denoting the characters relation/edge.
	CharactersColumn = "movie_characters"
	// QuotesTable is the table that holds the quotes relation/edge.
	QuotesTable = "movie_quotes"
	// QuotesInverseTable is the table name for the MovieQuote entity.
	// It exists in this package in order to avoid circular dependency with the "moviequote" package.
	QuotesInverseTable = "movie_quotes"
	// QuotesColumn is the table column denoting the quotes relation/edge.
	QuotesColumn = "movie_quotes"
)

// Columns holds all SQL columns for movie fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTitle,
	FieldYear,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "movies"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"category_movies",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// YearValidator is a validator for the "year" field. It is called by the builders before save.
	YearValidator func(int) error
)

// OrderOption defines the ordering options for the Movie queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByYear orders the results by the year field.
func ByYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldYear, opts...).ToFunc()
}

// ByCategoryField orders the results by category field.
func ByCategoryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCategoryStep(), sql.OrderByField(field, opts...))
	}
}

// ByCharactersCount orders the results by characters count.
func ByCharactersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCharactersStep(), opts...)
	}
}

// ByCharacters orders the results by characters terms.
func ByCharacters(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCharactersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByQuotesCount orders the results by quotes count.
func ByQuotesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newQuotesStep(), opts...)
	}
}

// ByQuotes orders the results by quotes terms.
func ByQuotes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newQuotesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCategoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
	)
}
func newCharactersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CharactersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CharactersTable, CharactersColumn),
	)
}
func newQuotesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(QuotesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, QuotesTable, QuotesColumn),
	)
}
