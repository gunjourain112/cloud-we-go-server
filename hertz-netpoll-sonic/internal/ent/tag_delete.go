package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/ent/predicate"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/ent/tag"
)

type TagDelete struct {
	config
	hooks    []Hook
	mutation *TagMutation
}

func (_d *TagDelete) Where(ps ...predicate.Tag) *TagDelete {
	_d.mutation.Where(ps...)
	return _d
}

func (_d *TagDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, _d.sqlExec, _d.mutation, _d.hooks)
}

func (_d *TagDelete) ExecX(ctx context.Context) int {
	n, err := _d.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (_d *TagDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(tag.Table, sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID))
	if ps := _d.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, _d.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	_d.mutation.done = true
	return affected, err
}

type TagDeleteOne struct {
	_d *TagDelete
}

func (_d *TagDeleteOne) Where(ps ...predicate.Tag) *TagDeleteOne {
	_d._d.mutation.Where(ps...)
	return _d
}

func (_d *TagDeleteOne) Exec(ctx context.Context) error {
	n, err := _d._d.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tag.Label}
	default:
		return nil
	}
}

func (_d *TagDeleteOne) ExecX(ctx context.Context) {
	if err := _d.Exec(ctx); err != nil {
		panic(err)
	}
}
