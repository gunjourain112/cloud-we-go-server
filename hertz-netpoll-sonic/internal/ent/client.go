package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/google/uuid"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/ent/post"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/ent/tag"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/ent/user"
)

type Client struct {
	config

	Schema *migrate.Schema

	Post *PostClient

	Tag *TagClient

	User *UserClient
}

func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Post = NewPostClient(c.config)
	c.Tag = NewTagClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	config struct {
		driver dialect.Driver

		debug bool

		log func(...any)

		hooks *hooks

		inters *inters
	}

	Option func(*config)
)

func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Post:   NewPostClient(cfg),
		Tag:    NewTagClient(cfg),
		User:   NewUserClient(cfg),
	}, nil
}

func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Post:   NewPostClient(cfg),
		Tag:    NewTagClient(cfg),
		User:   NewUserClient(cfg),
	}, nil
}

func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) Close() error {
	return c.driver.Close()
}

func (c *Client) Use(hooks ...Hook) {
	c.Post.Use(hooks...)
	c.Tag.Use(hooks...)
	c.User.Use(hooks...)
}

func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Post.Intercept(interceptors...)
	c.Tag.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *PostMutation:
		return c.Post.mutate(ctx, m)
	case *TagMutation:
		return c.Tag.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

type PostClient struct {
	config
}

func NewPostClient(c config) *PostClient {
	return &PostClient{config: c}
}

func (c *PostClient) Use(hooks ...Hook) {
	c.hooks.Post = append(c.hooks.Post, hooks...)
}

func (c *PostClient) Intercept(interceptors ...Interceptor) {
	c.inters.Post = append(c.inters.Post, interceptors...)
}

func (c *PostClient) Create() *PostCreate {
	mutation := newPostMutation(c.config, OpCreate)
	return &PostCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

func (c *PostClient) CreateBulk(builders ...*PostCreate) *PostCreateBulk {
	return &PostCreateBulk{config: c.config, builders: builders}
}

func (c *PostClient) MapCreateBulk(slice any, setFunc func(*PostCreate, int)) *PostCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PostCreateBulk{err: fmt.Errorf("calling to PostClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PostCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PostCreateBulk{config: c.config, builders: builders}
}

func (c *PostClient) Update() *PostUpdate {
	mutation := newPostMutation(c.config, OpUpdate)
	return &PostUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

func (c *PostClient) UpdateOne(_m *Post) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPost(_m))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

func (c *PostClient) UpdateOneID(id uuid.UUID) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPostID(id))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

func (c *PostClient) Delete() *PostDelete {
	mutation := newPostMutation(c.config, OpDelete)
	return &PostDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

func (c *PostClient) DeleteOne(_m *Post) *PostDeleteOne {
	return c.DeleteOneID(_m.ID)
}

func (c *PostClient) DeleteOneID(id uuid.UUID) *PostDeleteOne {
	builder := c.Delete().Where(post.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PostDeleteOne{builder}
}

func (c *PostClient) Query() *PostQuery {
	return &PostQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePost},
		inters: c.Interceptors(),
	}
}

func (c *PostClient) Get(ctx context.Context, id uuid.UUID) (*Post, error) {
	return c.Query().Where(post.ID(id)).Only(ctx)
}

func (c *PostClient) GetX(ctx context.Context, id uuid.UUID) *Post {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

func (c *PostClient) QueryAuthor(_m *Post) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := _m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, post.AuthorTable, post.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(_m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

func (c *PostClient) QueryTags(_m *Post) *TagQuery {
	query := (&TagClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := _m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, id),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, post.TagsTable, post.TagsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(_m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

func (c *PostClient) Hooks() []Hook {
	return c.hooks.Post
}

func (c *PostClient) Interceptors() []Interceptor {
	return c.inters.Post
}

func (c *PostClient) mutate(ctx context.Context, m *PostMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PostCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PostUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PostDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Post mutation op: %q", m.Op())
	}
}

type TagClient struct {
	config
}

func NewTagClient(c config) *TagClient {
	return &TagClient{config: c}
}

func (c *TagClient) Use(hooks ...Hook) {
	c.hooks.Tag = append(c.hooks.Tag, hooks...)
}

func (c *TagClient) Intercept(interceptors ...Interceptor) {
	c.inters.Tag = append(c.inters.Tag, interceptors...)
}

func (c *TagClient) Create() *TagCreate {
	mutation := newTagMutation(c.config, OpCreate)
	return &TagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

func (c *TagClient) CreateBulk(builders ...*TagCreate) *TagCreateBulk {
	return &TagCreateBulk{config: c.config, builders: builders}
}

func (c *TagClient) MapCreateBulk(slice any, setFunc func(*TagCreate, int)) *TagCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TagCreateBulk{err: fmt.Errorf("calling to TagClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TagCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TagCreateBulk{config: c.config, builders: builders}
}

func (c *TagClient) Update() *TagUpdate {
	mutation := newTagMutation(c.config, OpUpdate)
	return &TagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

func (c *TagClient) UpdateOne(_m *Tag) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTag(_m))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

func (c *TagClient) UpdateOneID(id uuid.UUID) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTagID(id))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

func (c *TagClient) Delete() *TagDelete {
	mutation := newTagMutation(c.config, OpDelete)
	return &TagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

func (c *TagClient) DeleteOne(_m *Tag) *TagDeleteOne {
	return c.DeleteOneID(_m.ID)
}

func (c *TagClient) DeleteOneID(id uuid.UUID) *TagDeleteOne {
	builder := c.Delete().Where(tag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TagDeleteOne{builder}
}

func (c *TagClient) Query() *TagQuery {
	return &TagQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTag},
		inters: c.Interceptors(),
	}
}

func (c *TagClient) Get(ctx context.Context, id uuid.UUID) (*Tag, error) {
	return c.Query().Where(tag.ID(id)).Only(ctx)
}

func (c *TagClient) GetX(ctx context.Context, id uuid.UUID) *Tag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

func (c *TagClient) QueryPosts(_m *Tag) *PostQuery {
	query := (&PostClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := _m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tag.Table, tag.FieldID, id),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tag.PostsTable, tag.PostsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(_m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

func (c *TagClient) Hooks() []Hook {
	return c.hooks.Tag
}

func (c *TagClient) Interceptors() []Interceptor {
	return c.inters.Tag
}

func (c *TagClient) mutate(ctx context.Context, m *TagMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TagCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TagUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TagDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Tag mutation op: %q", m.Op())
	}
}

type UserClient struct {
	config
}

func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

func (c *UserClient) UpdateOne(_m *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(_m))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

func (c *UserClient) DeleteOne(_m *User) *UserDeleteOne {
	return c.DeleteOneID(_m.ID)
}

func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

func (c *UserClient) QueryPosts(_m *User) *PostQuery {
	query := (&PostClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := _m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.PostsTable, user.PostsColumn),
		)
		fromV = sqlgraph.Neighbors(_m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

type (
	hooks struct {
		Post, Tag, User []ent.Hook
	}
	inters struct {
		Post, Tag, User []ent.Interceptor
	}
)
