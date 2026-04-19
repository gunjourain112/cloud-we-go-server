package comment

import (
	"time"
)

type CreateRequest struct {
	Body string `json:"body" validate:"required"`
}

type ReplyRequest struct {
	Body string `json:"body" validate:"required"`
}

type Response struct {
	ID        string    `json:"id"`
	PostID    string    `json:"post_id"`
	AuthorID  string    `json:"author_id"`
	Body      string    `json:"body"`
	Replies   []Reply   `json:"replies"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Reply struct {
	ID        string    `json:"id"`
	AuthorID  string    `json:"author_id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
