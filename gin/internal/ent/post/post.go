package post

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	Label = "post"

	FieldID = "id"

	FieldTitle = "title"

	FieldBody = "body"

	FieldAuthorID = "author_id"

	FieldCreatedAt = "created_at"

	FieldUpdatedAt = "updated_at"

	EdgeAuthor = "author"

	EdgeTags = "tags"

	Table = "posts"

	AuthorTable = "posts"

	AuthorInverseTable = "users"

	AuthorColumn = "author_id"

	TagsTable = "post_tags"

	TagsInverseTable = "tags"
)

var Columns = []string{
	FieldID,
	FieldTitle,
	FieldBody,
	FieldAuthorID,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	TagsPrimaryKey = []string{"post_id", "tag_id"}
)

func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	TitleValidator func(string) error

	BodyValidator func(string) error

	DefaultCreatedAt func() time.Time

	DefaultUpdatedAt func() time.Time

	UpdateDefaultUpdatedAt func() time.Time

	DefaultID func() uuid.UUID
)

type OrderOption func(*sql.Selector)

func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

func ByBody(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBody, opts...).ToFunc()
}

func ByAuthorID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthorID, opts...).ToFunc()
}

func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

func ByAuthorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorStep(), sql.OrderByField(field, opts...))
	}
}

func ByTagsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTagsStep(), opts...)
	}
}

func ByTags(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAuthorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
	)
}
func newTagsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
	)
}
