package predicate

import (
	"entgo.io/ent/dialect/sql"
)

type Post func(*sql.Selector)

type Tag func(*sql.Selector)

type User func(*sql.Selector)
