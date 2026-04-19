package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent"
)

type Repository interface {
	Create(ctx context.Context, email, passwordHash, name string) (*ent.User, error)
	GetByEmail(ctx context.Context, email string) (*ent.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.User, error)
}
