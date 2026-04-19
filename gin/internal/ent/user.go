package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent/user"
)

type User struct {
	config `json:"-"`

	ID uuid.UUID `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Email string `json:"email,omitempty"`

	PasswordHash string `json:"-"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

type UserEdges struct {
	Posts []*Post `json:"posts,omitempty"`

	loadedTypes [1]bool
}

func (e UserEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[0] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldName, user.FieldEmail, user.FieldPasswordHash:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

func (_m *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				_m.ID = *value
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				_m.Name = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				_m.Email = value.String
			}
		case user.FieldPasswordHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password_hash", values[i])
			} else if value.Valid {
				_m.PasswordHash = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				_m.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
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

func (_m *User) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

func (_m *User) QueryPosts() *PostQuery {
	return NewUserClient(_m.config).QueryPosts(_m)
}

func (_m *User) Update() *UserUpdateOne {
	return NewUserClient(_m.config).UpdateOne(_m)
}

func (_m *User) Unwrap() *User {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

func (_m *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("name=")
	builder.WriteString(_m.Name)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(_m.Email)
	builder.WriteString(", ")
	builder.WriteString("password_hash=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(_m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(_m.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

type Users []*User
