package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/ent/predicate"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/ent/user"
)

type UserDelete struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

func (_d *UserDelete) Where(ps ...predicate.User) *UserDelete {
	_d.mutation.Where(ps...)
	return _d
}

func (_d *UserDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, _d.sqlExec, _d.mutation, _d.hooks)
}

func (_d *UserDelete) ExecX(ctx context.Context) int {
	n, err := _d.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (_d *UserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
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

type UserDeleteOne struct {
	_d *UserDelete
}

func (_d *UserDeleteOne) Where(ps ...predicate.User) *UserDeleteOne {
	_d._d.mutation.Where(ps...)
	return _d
}

func (_d *UserDeleteOne) Exec(ctx context.Context) error {
	n, err := _d._d.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{user.Label}
	default:
		return nil
	}
}

func (_d *UserDeleteOne) ExecX(ctx context.Context) {
	if err := _d.Exec(ctx); err != nil {
		panic(err)
	}
}
