package tag

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent/predicate"
)

func ID(id uuid.UUID) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldID, id))
}

func IDEQ(id uuid.UUID) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldID, id))
}

func IDNEQ(id uuid.UUID) predicate.Tag {
	return predicate.Tag(sql.FieldNEQ(FieldID, id))
}

func IDIn(ids ...uuid.UUID) predicate.Tag {
	return predicate.Tag(sql.FieldIn(FieldID, ids...))
}

func IDNotIn(ids ...uuid.UUID) predicate.Tag {
	return predicate.Tag(sql.FieldNotIn(FieldID, ids...))
}

func IDGT(id uuid.UUID) predicate.Tag {
	return predicate.Tag(sql.FieldGT(FieldID, id))
}

func IDGTE(id uuid.UUID) predicate.Tag {
	return predicate.Tag(sql.FieldGTE(FieldID, id))
}

func IDLT(id uuid.UUID) predicate.Tag {
	return predicate.Tag(sql.FieldLT(FieldID, id))
}

func IDLTE(id uuid.UUID) predicate.Tag {
	return predicate.Tag(sql.FieldLTE(FieldID, id))
}

func Name(v string) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldName, v))
}

func NameEQ(v string) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldName, v))
}

func NameNEQ(v string) predicate.Tag {
	return predicate.Tag(sql.FieldNEQ(FieldName, v))
}

func NameIn(vs ...string) predicate.Tag {
	return predicate.Tag(sql.FieldIn(FieldName, vs...))
}

func NameNotIn(vs ...string) predicate.Tag {
	return predicate.Tag(sql.FieldNotIn(FieldName, vs...))
}

func NameGT(v string) predicate.Tag {
	return predicate.Tag(sql.FieldGT(FieldName, v))
}

func NameGTE(v string) predicate.Tag {
	return predicate.Tag(sql.FieldGTE(FieldName, v))
}

func NameLT(v string) predicate.Tag {
	return predicate.Tag(sql.FieldLT(FieldName, v))
}

func NameLTE(v string) predicate.Tag {
	return predicate.Tag(sql.FieldLTE(FieldName, v))
}

func NameContains(v string) predicate.Tag {
	return predicate.Tag(sql.FieldContains(FieldName, v))
}

func NameHasPrefix(v string) predicate.Tag {
	return predicate.Tag(sql.FieldHasPrefix(FieldName, v))
}

func NameHasSuffix(v string) predicate.Tag {
	return predicate.Tag(sql.FieldHasSuffix(FieldName, v))
}

func NameEqualFold(v string) predicate.Tag {
	return predicate.Tag(sql.FieldEqualFold(FieldName, v))
}

func NameContainsFold(v string) predicate.Tag {
	return predicate.Tag(sql.FieldContainsFold(FieldName, v))
}

func HasPosts() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PostsTable, PostsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

func HasPostsWith(preds ...predicate.Post) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		step := newPostsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

func And(predicates ...predicate.Tag) predicate.Tag {
	return predicate.Tag(sql.AndPredicates(predicates...))
}

func Or(predicates ...predicate.Tag) predicate.Tag {
	return predicate.Tag(sql.OrPredicates(predicates...))
}

func Not(p predicate.Tag) predicate.Tag {
	return predicate.Tag(sql.NotPredicates(p))
}
