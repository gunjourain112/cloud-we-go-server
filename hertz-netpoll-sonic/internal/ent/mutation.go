package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/ent/post"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/ent/predicate"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/ent/tag"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/ent/user"
)

const (
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	TypePost = "Post"
	TypeTag  = "Tag"
	TypeUser = "User"
)

type PostMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	title         *string
	body          *string
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	author        *uuid.UUID
	clearedauthor bool
	tags          map[uuid.UUID]struct{}
	removedtags   map[uuid.UUID]struct{}
	clearedtags   bool
	done          bool
	oldValue      func(context.Context) (*Post, error)
	predicates    []predicate.Post
}

var _ ent.Mutation = (*PostMutation)(nil)

type postOption func(*PostMutation)

func newPostMutation(c config, op Op, opts ...postOption) *PostMutation {
	m := &PostMutation{
		config:        c,
		op:            op,
		typ:           TypePost,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func withPostID(id uuid.UUID) postOption {
	return func(m *PostMutation) {
		var (
			err   error
			once  sync.Once
			value *Post
		)
		m.oldValue = func(ctx context.Context) (*Post, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Post.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

func withPost(node *Post) postOption {
	return func(m *PostMutation) {
		m.oldValue = func(context.Context) (*Post, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

func (m PostMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

func (m PostMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

func (m *PostMutation) SetID(id uuid.UUID) {
	m.id = &id
}

func (m *PostMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

func (m *PostMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Post.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

func (m *PostMutation) SetTitle(s string) {
	m.title = &s
}

func (m *PostMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

func (m *PostMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

func (m *PostMutation) ResetTitle() {
	m.title = nil
}

func (m *PostMutation) SetBody(s string) {
	m.body = &s
}

func (m *PostMutation) Body() (r string, exists bool) {
	v := m.body
	if v == nil {
		return
	}
	return *v, true
}

func (m *PostMutation) OldBody(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBody is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBody requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBody: %w", err)
	}
	return oldValue.Body, nil
}

func (m *PostMutation) ResetBody() {
	m.body = nil
}

func (m *PostMutation) SetAuthorID(u uuid.UUID) {
	m.author = &u
}

func (m *PostMutation) AuthorID() (r uuid.UUID, exists bool) {
	v := m.author
	if v == nil {
		return
	}
	return *v, true
}

func (m *PostMutation) OldAuthorID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAuthorID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAuthorID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAuthorID: %w", err)
	}
	return oldValue.AuthorID, nil
}

func (m *PostMutation) ResetAuthorID() {
	m.author = nil
}

func (m *PostMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

func (m *PostMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

func (m *PostMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

func (m *PostMutation) ResetCreatedAt() {
	m.created_at = nil
}

func (m *PostMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

func (m *PostMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

func (m *PostMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

func (m *PostMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

func (m *PostMutation) ClearAuthor() {
	m.clearedauthor = true
	m.clearedFields[post.FieldAuthorID] = struct{}{}
}

func (m *PostMutation) AuthorCleared() bool {
	return m.clearedauthor
}

func (m *PostMutation) AuthorIDs() (ids []uuid.UUID) {
	if id := m.author; id != nil {
		ids = append(ids, *id)
	}
	return
}

func (m *PostMutation) ResetAuthor() {
	m.author = nil
	m.clearedauthor = false
}

func (m *PostMutation) AddTagIDs(ids ...uuid.UUID) {
	if m.tags == nil {
		m.tags = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.tags[ids[i]] = struct{}{}
	}
}

func (m *PostMutation) ClearTags() {
	m.clearedtags = true
}

func (m *PostMutation) TagsCleared() bool {
	return m.clearedtags
}

func (m *PostMutation) RemoveTagIDs(ids ...uuid.UUID) {
	if m.removedtags == nil {
		m.removedtags = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		delete(m.tags, ids[i])
		m.removedtags[ids[i]] = struct{}{}
	}
}

func (m *PostMutation) RemovedTagsIDs() (ids []uuid.UUID) {
	for id := range m.removedtags {
		ids = append(ids, id)
	}
	return
}

func (m *PostMutation) TagsIDs() (ids []uuid.UUID) {
	for id := range m.tags {
		ids = append(ids, id)
	}
	return
}

func (m *PostMutation) ResetTags() {
	m.tags = nil
	m.clearedtags = false
	m.removedtags = nil
}

func (m *PostMutation) Where(ps ...predicate.Post) {
	m.predicates = append(m.predicates, ps...)
}

func (m *PostMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Post, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

func (m *PostMutation) Op() Op {
	return m.op
}

func (m *PostMutation) SetOp(op Op) {
	m.op = op
}

func (m *PostMutation) Type() string {
	return m.typ
}

func (m *PostMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.title != nil {
		fields = append(fields, post.FieldTitle)
	}
	if m.body != nil {
		fields = append(fields, post.FieldBody)
	}
	if m.author != nil {
		fields = append(fields, post.FieldAuthorID)
	}
	if m.created_at != nil {
		fields = append(fields, post.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, post.FieldUpdatedAt)
	}
	return fields
}

func (m *PostMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case post.FieldTitle:
		return m.Title()
	case post.FieldBody:
		return m.Body()
	case post.FieldAuthorID:
		return m.AuthorID()
	case post.FieldCreatedAt:
		return m.CreatedAt()
	case post.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

func (m *PostMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case post.FieldTitle:
		return m.OldTitle(ctx)
	case post.FieldBody:
		return m.OldBody(ctx)
	case post.FieldAuthorID:
		return m.OldAuthorID(ctx)
	case post.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case post.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Post field %s", name)
}

func (m *PostMutation) SetField(name string, value ent.Value) error {
	switch name {
	case post.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case post.FieldBody:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBody(v)
		return nil
	case post.FieldAuthorID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAuthorID(v)
		return nil
	case post.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case post.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Post field %s", name)
}

func (m *PostMutation) AddedFields() []string {
	return nil
}

func (m *PostMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

func (m *PostMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Post numeric field %s", name)
}

func (m *PostMutation) ClearedFields() []string {
	return nil
}

func (m *PostMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

func (m *PostMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Post nullable field %s", name)
}

func (m *PostMutation) ResetField(name string) error {
	switch name {
	case post.FieldTitle:
		m.ResetTitle()
		return nil
	case post.FieldBody:
		m.ResetBody()
		return nil
	case post.FieldAuthorID:
		m.ResetAuthorID()
		return nil
	case post.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case post.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Post field %s", name)
}

func (m *PostMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.author != nil {
		edges = append(edges, post.EdgeAuthor)
	}
	if m.tags != nil {
		edges = append(edges, post.EdgeTags)
	}
	return edges
}

func (m *PostMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case post.EdgeAuthor:
		if id := m.author; id != nil {
			return []ent.Value{*id}
		}
	case post.EdgeTags:
		ids := make([]ent.Value, 0, len(m.tags))
		for id := range m.tags {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

func (m *PostMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedtags != nil {
		edges = append(edges, post.EdgeTags)
	}
	return edges
}

func (m *PostMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case post.EdgeTags:
		ids := make([]ent.Value, 0, len(m.removedtags))
		for id := range m.removedtags {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

func (m *PostMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedauthor {
		edges = append(edges, post.EdgeAuthor)
	}
	if m.clearedtags {
		edges = append(edges, post.EdgeTags)
	}
	return edges
}

func (m *PostMutation) EdgeCleared(name string) bool {
	switch name {
	case post.EdgeAuthor:
		return m.clearedauthor
	case post.EdgeTags:
		return m.clearedtags
	}
	return false
}

func (m *PostMutation) ClearEdge(name string) error {
	switch name {
	case post.EdgeAuthor:
		m.ClearAuthor()
		return nil
	}
	return fmt.Errorf("unknown Post unique edge %s", name)
}

func (m *PostMutation) ResetEdge(name string) error {
	switch name {
	case post.EdgeAuthor:
		m.ResetAuthor()
		return nil
	case post.EdgeTags:
		m.ResetTags()
		return nil
	}
	return fmt.Errorf("unknown Post edge %s", name)
}

type TagMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	name          *string
	clearedFields map[string]struct{}
	posts         map[uuid.UUID]struct{}
	removedposts  map[uuid.UUID]struct{}
	clearedposts  bool
	done          bool
	oldValue      func(context.Context) (*Tag, error)
	predicates    []predicate.Tag
}

var _ ent.Mutation = (*TagMutation)(nil)

type tagOption func(*TagMutation)

func newTagMutation(c config, op Op, opts ...tagOption) *TagMutation {
	m := &TagMutation{
		config:        c,
		op:            op,
		typ:           TypeTag,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func withTagID(id uuid.UUID) tagOption {
	return func(m *TagMutation) {
		var (
			err   error
			once  sync.Once
			value *Tag
		)
		m.oldValue = func(ctx context.Context) (*Tag, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Tag.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

func withTag(node *Tag) tagOption {
	return func(m *TagMutation) {
		m.oldValue = func(context.Context) (*Tag, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

func (m TagMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

func (m TagMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

func (m *TagMutation) SetID(id uuid.UUID) {
	m.id = &id
}

func (m *TagMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

func (m *TagMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Tag.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

func (m *TagMutation) SetName(s string) {
	m.name = &s
}

func (m *TagMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

func (m *TagMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

func (m *TagMutation) ResetName() {
	m.name = nil
}

func (m *TagMutation) AddPostIDs(ids ...uuid.UUID) {
	if m.posts == nil {
		m.posts = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.posts[ids[i]] = struct{}{}
	}
}

func (m *TagMutation) ClearPosts() {
	m.clearedposts = true
}

func (m *TagMutation) PostsCleared() bool {
	return m.clearedposts
}

func (m *TagMutation) RemovePostIDs(ids ...uuid.UUID) {
	if m.removedposts == nil {
		m.removedposts = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		delete(m.posts, ids[i])
		m.removedposts[ids[i]] = struct{}{}
	}
}

func (m *TagMutation) RemovedPostsIDs() (ids []uuid.UUID) {
	for id := range m.removedposts {
		ids = append(ids, id)
	}
	return
}

func (m *TagMutation) PostsIDs() (ids []uuid.UUID) {
	for id := range m.posts {
		ids = append(ids, id)
	}
	return
}

func (m *TagMutation) ResetPosts() {
	m.posts = nil
	m.clearedposts = false
	m.removedposts = nil
}

func (m *TagMutation) Where(ps ...predicate.Tag) {
	m.predicates = append(m.predicates, ps...)
}

func (m *TagMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Tag, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

func (m *TagMutation) Op() Op {
	return m.op
}

func (m *TagMutation) SetOp(op Op) {
	m.op = op
}

func (m *TagMutation) Type() string {
	return m.typ
}

func (m *TagMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, tag.FieldName)
	}
	return fields
}

func (m *TagMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case tag.FieldName:
		return m.Name()
	}
	return nil, false
}

func (m *TagMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case tag.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Tag field %s", name)
}

func (m *TagMutation) SetField(name string, value ent.Value) error {
	switch name {
	case tag.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Tag field %s", name)
}

func (m *TagMutation) AddedFields() []string {
	return nil
}

func (m *TagMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

func (m *TagMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Tag numeric field %s", name)
}

func (m *TagMutation) ClearedFields() []string {
	return nil
}

func (m *TagMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

func (m *TagMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Tag nullable field %s", name)
}

func (m *TagMutation) ResetField(name string) error {
	switch name {
	case tag.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Tag field %s", name)
}

func (m *TagMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.posts != nil {
		edges = append(edges, tag.EdgePosts)
	}
	return edges
}

func (m *TagMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case tag.EdgePosts:
		ids := make([]ent.Value, 0, len(m.posts))
		for id := range m.posts {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

func (m *TagMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedposts != nil {
		edges = append(edges, tag.EdgePosts)
	}
	return edges
}

func (m *TagMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case tag.EdgePosts:
		ids := make([]ent.Value, 0, len(m.removedposts))
		for id := range m.removedposts {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

func (m *TagMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedposts {
		edges = append(edges, tag.EdgePosts)
	}
	return edges
}

func (m *TagMutation) EdgeCleared(name string) bool {
	switch name {
	case tag.EdgePosts:
		return m.clearedposts
	}
	return false
}

func (m *TagMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Tag unique edge %s", name)
}

func (m *TagMutation) ResetEdge(name string) error {
	switch name {
	case tag.EdgePosts:
		m.ResetPosts()
		return nil
	}
	return fmt.Errorf("unknown Tag edge %s", name)
}

type UserMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	name          *string
	email         *string
	password_hash *string
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	posts         map[uuid.UUID]struct{}
	removedposts  map[uuid.UUID]struct{}
	clearedposts  bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

type userOption func(*UserMutation)

func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func withUserID(id uuid.UUID) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

func (m *UserMutation) SetID(id uuid.UUID) {
	m.id = &id
}

func (m *UserMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

func (m *UserMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

func (m *UserMutation) SetName(s string) {
	m.name = &s
}

func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

func (m *UserMutation) ResetName() {
	m.name = nil
}

func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

func (m *UserMutation) ResetEmail() {
	m.email = nil
}

func (m *UserMutation) SetPasswordHash(s string) {
	m.password_hash = &s
}

func (m *UserMutation) PasswordHash() (r string, exists bool) {
	v := m.password_hash
	if v == nil {
		return
	}
	return *v, true
}

func (m *UserMutation) OldPasswordHash(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPasswordHash is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPasswordHash requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPasswordHash: %w", err)
	}
	return oldValue.PasswordHash, nil
}

func (m *UserMutation) ResetPasswordHash() {
	m.password_hash = nil
}

func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

func (m *UserMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

func (m *UserMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

func (m *UserMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

func (m *UserMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

func (m *UserMutation) AddPostIDs(ids ...uuid.UUID) {
	if m.posts == nil {
		m.posts = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.posts[ids[i]] = struct{}{}
	}
}

func (m *UserMutation) ClearPosts() {
	m.clearedposts = true
}

func (m *UserMutation) PostsCleared() bool {
	return m.clearedposts
}

func (m *UserMutation) RemovePostIDs(ids ...uuid.UUID) {
	if m.removedposts == nil {
		m.removedposts = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		delete(m.posts, ids[i])
		m.removedposts[ids[i]] = struct{}{}
	}
}

func (m *UserMutation) RemovedPostsIDs() (ids []uuid.UUID) {
	for id := range m.removedposts {
		ids = append(ids, id)
	}
	return
}

func (m *UserMutation) PostsIDs() (ids []uuid.UUID) {
	for id := range m.posts {
		ids = append(ids, id)
	}
	return
}

func (m *UserMutation) ResetPosts() {
	m.posts = nil
	m.clearedposts = false
	m.removedposts = nil
}

func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

func (m *UserMutation) Op() Op {
	return m.op
}

func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

func (m *UserMutation) Type() string {
	return m.typ
}

func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.password_hash != nil {
		fields = append(fields, user.FieldPasswordHash)
	}
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, user.FieldUpdatedAt)
	}
	return fields
}

func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	case user.FieldEmail:
		return m.Email()
	case user.FieldPasswordHash:
		return m.PasswordHash()
	case user.FieldCreatedAt:
		return m.CreatedAt()
	case user.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldPasswordHash:
		return m.OldPasswordHash(ctx)
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case user.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldPasswordHash:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPasswordHash(v)
		return nil
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case user.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

func (m *UserMutation) AddedFields() []string {
	return nil
}

func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

func (m *UserMutation) ClearedFields() []string {
	return nil
}

func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldPasswordHash:
		m.ResetPasswordHash()
		return nil
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case user.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.posts != nil {
		edges = append(edges, user.EdgePosts)
	}
	return edges
}

func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgePosts:
		ids := make([]ent.Value, 0, len(m.posts))
		for id := range m.posts {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedposts != nil {
		edges = append(edges, user.EdgePosts)
	}
	return edges
}

func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgePosts:
		ids := make([]ent.Value, 0, len(m.removedposts))
		for id := range m.removedposts {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedposts {
		edges = append(edges, user.EdgePosts)
	}
	return edges
}

func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgePosts:
		return m.clearedposts
	}
	return false
}

func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgePosts:
		m.ResetPosts()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
