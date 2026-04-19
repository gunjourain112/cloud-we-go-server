package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent"
)

type Repository interface {
	Create(ctx context.Context, title, body string, authorID uuid.UUID, tags []string) (*ent.Post, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Post, error)
	List(ctx context.Context, limit, offset int) ([]*ent.Post, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
