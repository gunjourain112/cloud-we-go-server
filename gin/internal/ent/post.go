package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent/post"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent/user"
)

type Post struct {
	config `json:"-"`

	ID uuid.UUID `json:"id,omitempty"`

	Title string `json:"title,omitempty"`

	Body string `json:"body,omitempty"`

	AuthorID uuid.UUID `json:"author_id,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Edges        PostEdges `json:"edges"`
	selectValues sql.SelectValues
}

type PostEdges struct {
	Author *User `json:"author,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`

	loadedTypes [2]bool
}

func (e PostEdges) AuthorOrErr() (*User, error) {
	if e.Author != nil {
		return e.Author, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "author"}
}

func (e PostEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[1] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

func (*Post) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case post.FieldTitle, post.FieldBody:
			values[i] = new(sql.NullString)
		case post.FieldCreatedAt, post.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case post.FieldID, post.FieldAuthorID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

func (_m *Post) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case post.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				_m.ID = *value
			}
		case post.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				_m.Title = value.String
			}
		case post.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				_m.Body = value.String
			}
		case post.FieldAuthorID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field author_id", values[i])
			} else if value != nil {
				_m.AuthorID = *value
			}
		case post.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				_m.CreatedAt = value.Time
			}
		case post.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				_m.UpdatedAt = value.Time
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

func (_m *Post) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

func (_m *Post) QueryAuthor() *UserQuery {
	return NewPostClient(_m.config).QueryAuthor(_m)
}

func (_m *Post) QueryTags() *TagQuery {
	return NewPostClient(_m.config).QueryTags(_m)
}

func (_m *Post) Update() *PostUpdateOne {
	return NewPostClient(_m.config).UpdateOne(_m)
}

func (_m *Post) Unwrap() *Post {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Post is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

func (_m *Post) String() string {
	var builder strings.Builder
	builder.WriteString("Post(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("title=")
	builder.WriteString(_m.Title)
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(_m.Body)
	builder.WriteString(", ")
	builder.WriteString("author_id=")
	builder.WriteString(fmt.Sprintf("%v", _m.AuthorID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(_m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(_m.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

type Posts []*Post
