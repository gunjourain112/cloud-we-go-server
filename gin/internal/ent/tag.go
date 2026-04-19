package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent/tag"
)

type Tag struct {
	config `json:"-"`

	ID uuid.UUID `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Edges        TagEdges `json:"edges"`
	selectValues sql.SelectValues
}

type TagEdges struct {
	Posts []*Post `json:"posts,omitempty"`

	loadedTypes [1]bool
}

func (e TagEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[0] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

func (*Tag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tag.FieldName:
			values[i] = new(sql.NullString)
		case tag.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

func (_m *Tag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tag.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				_m.ID = *value
			}
		case tag.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				_m.Name = value.String
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

func (_m *Tag) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

func (_m *Tag) QueryPosts() *PostQuery {
	return NewTagClient(_m.config).QueryPosts(_m)
}

func (_m *Tag) Update() *TagUpdateOne {
	return NewTagClient(_m.config).UpdateOne(_m)
}

func (_m *Tag) Unwrap() *Tag {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tag is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

func (_m *Tag) String() string {
	var builder strings.Builder
	builder.WriteString("Tag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("name=")
	builder.WriteString(_m.Name)
	builder.WriteByte(')')
	return builder.String()
}

type Tags []*Tag
