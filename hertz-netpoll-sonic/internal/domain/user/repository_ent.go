package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/ent"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/ent/user"
)

type entRepository struct {
	client *ent.Client
}

func NewRepository(client *ent.Client) Repository {
	return &entRepository{client: client}
}

func (r *entRepository) Create(ctx context.Context, email, passwordHash, name string) (*ent.User, error) {
	return r.client.User.Create().
		SetEmail(email).
		SetPasswordHash(passwordHash).
		SetName(name).
		Save(ctx)
}

func (r *entRepository) GetByEmail(ctx context.Context, email string) (*ent.User, error) {
	return r.client.User.Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
}

func (r *entRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return r.client.User.Get(ctx, id)
}
