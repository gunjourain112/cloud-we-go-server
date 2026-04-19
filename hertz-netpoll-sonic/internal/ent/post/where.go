package post

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/ent/predicate"
)

func ID(id uuid.UUID) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

func IDEQ(id uuid.UUID) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

func IDNEQ(id uuid.UUID) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldID, id))
}

func IDIn(ids ...uuid.UUID) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldID, ids...))
}

func IDNotIn(ids ...uuid.UUID) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldID, ids...))
}

func IDGT(id uuid.UUID) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldID, id))
}

func IDGTE(id uuid.UUID) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldID, id))
}

func IDLT(id uuid.UUID) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldID, id))
}

func IDLTE(id uuid.UUID) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldID, id))
}

func Title(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldTitle, v))
}

func Body(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldBody, v))
}

func AuthorID(v uuid.UUID) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldAuthorID, v))
}

func CreatedAt(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreatedAt, v))
}

func UpdatedAt(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUpdatedAt, v))
}

func TitleEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldTitle, v))
}

func TitleNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldTitle, v))
}

func TitleIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldTitle, vs...))
}

func TitleNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldTitle, vs...))
}

func TitleGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldTitle, v))
}

func TitleGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldTitle, v))
}

func TitleLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldTitle, v))
}

func TitleLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldTitle, v))
}

func TitleContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldTitle, v))
}

func TitleHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldTitle, v))
}

func TitleHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldTitle, v))
}

func TitleEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldTitle, v))
}

func TitleContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldTitle, v))
}

func BodyEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldBody, v))
}

func BodyNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldBody, v))
}

func BodyIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldBody, vs...))
}

func BodyNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldBody, vs...))
}

func BodyGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldBody, v))
}

func BodyGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldBody, v))
}

func BodyLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldBody, v))
}

func BodyLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldBody, v))
}

func BodyContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldBody, v))
}

func BodyHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldBody, v))
}

func BodyHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldBody, v))
}

func BodyEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldBody, v))
}

func BodyContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldBody, v))
}

func AuthorIDEQ(v uuid.UUID) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldAuthorID, v))
}

func AuthorIDNEQ(v uuid.UUID) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldAuthorID, v))
}

func AuthorIDIn(vs ...uuid.UUID) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldAuthorID, vs...))
}

func AuthorIDNotIn(vs ...uuid.UUID) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldAuthorID, vs...))
}

func CreatedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreatedAt, v))
}

func CreatedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldCreatedAt, v))
}

func CreatedAtIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldCreatedAt, vs...))
}

func CreatedAtNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldCreatedAt, vs...))
}

func CreatedAtGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldCreatedAt, v))
}

func CreatedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldCreatedAt, v))
}

func CreatedAtLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldCreatedAt, v))
}

func CreatedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldCreatedAt, v))
}

func UpdatedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUpdatedAt, v))
}

func UpdatedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldUpdatedAt, v))
}

func UpdatedAtIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldUpdatedAt, vs...))
}

func UpdatedAtNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

func UpdatedAtGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldUpdatedAt, v))
}

func UpdatedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldUpdatedAt, v))
}

func UpdatedAtLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldUpdatedAt, v))
}

func UpdatedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldUpdatedAt, v))
}

func HasAuthor() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

func HasAuthorWith(preds ...predicate.User) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := newAuthorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

func HasTags() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

func HasTagsWith(preds ...predicate.Tag) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := newTagsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

func And(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(sql.AndPredicates(predicates...))
}

func Or(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(sql.OrPredicates(predicates...))
}

func Not(p predicate.Post) predicate.Post {
	return predicate.Post(sql.NotPredicates(p))
}
