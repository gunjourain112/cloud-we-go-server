package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent/post"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent/predicate"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent/tag"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent/user"
)

type PostUpdate struct {
	config
	hooks    []Hook
	mutation *PostMutation
}

func (_u *PostUpdate) Where(ps ...predicate.Post) *PostUpdate {
	_u.mutation.Where(ps...)
	return _u
}

func (_u *PostUpdate) SetTitle(v string) *PostUpdate {
	_u.mutation.SetTitle(v)
	return _u
}

func (_u *PostUpdate) SetNillableTitle(v *string) *PostUpdate {
	if v != nil {
		_u.SetTitle(*v)
	}
	return _u
}

func (_u *PostUpdate) SetBody(v string) *PostUpdate {
	_u.mutation.SetBody(v)
	return _u
}

func (_u *PostUpdate) SetNillableBody(v *string) *PostUpdate {
	if v != nil {
		_u.SetBody(*v)
	}
	return _u
}

func (_u *PostUpdate) SetAuthorID(v uuid.UUID) *PostUpdate {
	_u.mutation.SetAuthorID(v)
	return _u
}

func (_u *PostUpdate) SetNillableAuthorID(v *uuid.UUID) *PostUpdate {
	if v != nil {
		_u.SetAuthorID(*v)
	}
	return _u
}

func (_u *PostUpdate) SetUpdatedAt(v time.Time) *PostUpdate {
	_u.mutation.SetUpdatedAt(v)
	return _u
}

func (_u *PostUpdate) SetAuthor(v *User) *PostUpdate {
	return _u.SetAuthorID(v.ID)
}

func (_u *PostUpdate) AddTagIDs(ids ...uuid.UUID) *PostUpdate {
	_u.mutation.AddTagIDs(ids...)
	return _u
}

func (_u *PostUpdate) AddTags(v ...*Tag) *PostUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddTagIDs(ids...)
}

func (_u *PostUpdate) Mutation() *PostMutation {
	return _u.mutation
}

func (_u *PostUpdate) ClearAuthor() *PostUpdate {
	_u.mutation.ClearAuthor()
	return _u
}

func (_u *PostUpdate) ClearTags() *PostUpdate {
	_u.mutation.ClearTags()
	return _u
}

func (_u *PostUpdate) RemoveTagIDs(ids ...uuid.UUID) *PostUpdate {
	_u.mutation.RemoveTagIDs(ids...)
	return _u
}

func (_u *PostUpdate) RemoveTags(v ...*Tag) *PostUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveTagIDs(ids...)
}

func (_u *PostUpdate) Save(ctx context.Context) (int, error) {
	_u.defaults()
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

func (_u *PostUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

func (_u *PostUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

func (_u *PostUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *PostUpdate) defaults() {
	if _, ok := _u.mutation.UpdatedAt(); !ok {
		v := post.UpdateDefaultUpdatedAt()
		_u.mutation.SetUpdatedAt(v)
	}
}

func (_u *PostUpdate) check() error {
	if v, ok := _u.mutation.Title(); ok {
		if err := post.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Post.title": %w`, err)}
		}
	}
	if v, ok := _u.mutation.Body(); ok {
		if err := post.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "Post.body": %w`, err)}
		}
	}
	if _u.mutation.AuthorCleared() && len(_u.mutation.AuthorIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Post.author"`)
	}
	return nil
}

func (_u *PostUpdate) sqlSave(ctx context.Context) (_node int, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeUUID))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Title(); ok {
		_spec.SetField(post.FieldTitle, field.TypeString, value)
	}
	if value, ok := _u.mutation.Body(); ok {
		_spec.SetField(post.FieldBody, field.TypeString, value)
	}
	if value, ok := _u.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
	}
	if _u.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.AuthorTable,
			Columns: []string{post.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.AuthorTable,
			Columns: []string{post.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.TagsTable,
			Columns: post.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedTagsIDs(); len(nodes) > 0 && !_u.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.TagsTable,
			Columns: post.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.TagsTable,
			Columns: post.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _node, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return _node, nil
}

type PostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PostMutation
}

func (_u *PostUpdateOne) SetTitle(v string) *PostUpdateOne {
	_u.mutation.SetTitle(v)
	return _u
}

func (_u *PostUpdateOne) SetNillableTitle(v *string) *PostUpdateOne {
	if v != nil {
		_u.SetTitle(*v)
	}
	return _u
}

func (_u *PostUpdateOne) SetBody(v string) *PostUpdateOne {
	_u.mutation.SetBody(v)
	return _u
}

func (_u *PostUpdateOne) SetNillableBody(v *string) *PostUpdateOne {
	if v != nil {
		_u.SetBody(*v)
	}
	return _u
}

func (_u *PostUpdateOne) SetAuthorID(v uuid.UUID) *PostUpdateOne {
	_u.mutation.SetAuthorID(v)
	return _u
}

func (_u *PostUpdateOne) SetNillableAuthorID(v *uuid.UUID) *PostUpdateOne {
	if v != nil {
		_u.SetAuthorID(*v)
	}
	return _u
}

func (_u *PostUpdateOne) SetUpdatedAt(v time.Time) *PostUpdateOne {
	_u.mutation.SetUpdatedAt(v)
	return _u
}

func (_u *PostUpdateOne) SetAuthor(v *User) *PostUpdateOne {
	return _u.SetAuthorID(v.ID)
}

func (_u *PostUpdateOne) AddTagIDs(ids ...uuid.UUID) *PostUpdateOne {
	_u.mutation.AddTagIDs(ids...)
	return _u
}

func (_u *PostUpdateOne) AddTags(v ...*Tag) *PostUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddTagIDs(ids...)
}

func (_u *PostUpdateOne) Mutation() *PostMutation {
	return _u.mutation
}

func (_u *PostUpdateOne) ClearAuthor() *PostUpdateOne {
	_u.mutation.ClearAuthor()
	return _u
}

func (_u *PostUpdateOne) ClearTags() *PostUpdateOne {
	_u.mutation.ClearTags()
	return _u
}

func (_u *PostUpdateOne) RemoveTagIDs(ids ...uuid.UUID) *PostUpdateOne {
	_u.mutation.RemoveTagIDs(ids...)
	return _u
}

func (_u *PostUpdateOne) RemoveTags(v ...*Tag) *PostUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveTagIDs(ids...)
}

func (_u *PostUpdateOne) Where(ps ...predicate.Post) *PostUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

func (_u *PostUpdateOne) Select(field string, fields ...string) *PostUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

func (_u *PostUpdateOne) Save(ctx context.Context) (*Post, error) {
	_u.defaults()
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

func (_u *PostUpdateOne) SaveX(ctx context.Context) *Post {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

func (_u *PostUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

func (_u *PostUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *PostUpdateOne) defaults() {
	if _, ok := _u.mutation.UpdatedAt(); !ok {
		v := post.UpdateDefaultUpdatedAt()
		_u.mutation.SetUpdatedAt(v)
	}
}

func (_u *PostUpdateOne) check() error {
	if v, ok := _u.mutation.Title(); ok {
		if err := post.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Post.title": %w`, err)}
		}
	}
	if v, ok := _u.mutation.Body(); ok {
		if err := post.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "Post.body": %w`, err)}
		}
	}
	if _u.mutation.AuthorCleared() && len(_u.mutation.AuthorIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Post.author"`)
	}
	return nil
}

func (_u *PostUpdateOne) sqlSave(ctx context.Context) (_node *Post, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeUUID))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Post.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, post.FieldID)
		for _, f := range fields {
			if !post.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != post.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Title(); ok {
		_spec.SetField(post.FieldTitle, field.TypeString, value)
	}
	if value, ok := _u.mutation.Body(); ok {
		_spec.SetField(post.FieldBody, field.TypeString, value)
	}
	if value, ok := _u.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
	}
	if _u.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.AuthorTable,
			Columns: []string{post.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.AuthorTable,
			Columns: []string{post.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.TagsTable,
			Columns: post.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedTagsIDs(); len(nodes) > 0 && !_u.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.TagsTable,
			Columns: post.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.TagsTable,
			Columns: post.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Post{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
