package post

import (
	"time"

	"github.com/google/uuid"
)

type CreateRequest struct {
	Title string   `json:"title" validate:"required"`
	Body  string   `json:"body" validate:"required"`
	Tags  []string `json:"tags"`
}

type Response struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	AuthorID  uuid.UUID `json:"author_id"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
