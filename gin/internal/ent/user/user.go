package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	Label = "user"

	FieldID = "id"

	FieldName = "name"

	FieldEmail = "email"

	FieldPasswordHash = "password_hash"

	FieldCreatedAt = "created_at"

	FieldUpdatedAt = "updated_at"

	EdgePosts = "posts"

	Table = "users"

	PostsTable = "posts"

	PostsInverseTable = "posts"

	PostsColumn = "author_id"
)

var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldPasswordHash,
	FieldCreatedAt,
	FieldUpdatedAt,
}

func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	NameValidator func(string) error

	EmailValidator func(string) error

	DefaultCreatedAt func() time.Time

	DefaultUpdatedAt func() time.Time

	UpdateDefaultUpdatedAt func() time.Time

	DefaultID func() uuid.UUID
)

type OrderOption func(*sql.Selector)

func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

func ByPasswordHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPasswordHash, opts...).ToFunc()
}

func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

func ByPostsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPostsStep(), opts...)
	}
}

func ByPosts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPostsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PostsTable, PostsColumn),
	)
}
