package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent/post"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent/tag"
)

type entRepository struct {
	client *ent.Client
}

func NewRepository(client *ent.Client) Repository {
	return &entRepository{client: client}
}

func (r *entRepository) Create(ctx context.Context, title, body string, authorID uuid.UUID, tags []string) (*ent.Post, error) {
	var tagIDs []uuid.UUID
	for _, t := range tags {
		entTag, err := r.client.Tag.Query().Where(tag.NameEQ(t)).Only(ctx)
		if err != nil {
			entTag, err = r.client.Tag.Create().SetName(t).Save(ctx)
			if err != nil {
				return nil, err
			}
		}
		tagIDs = append(tagIDs, entTag.ID)
	}

	return r.client.Post.Create().
		SetTitle(title).
		SetBody(body).
		SetAuthorID(authorID).
		AddTagIDs(tagIDs...).
		Save(ctx)
}

func (r *entRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.Post, error) {
	return r.client.Post.Query().
		Where(post.IDEQ(id)).
		WithTags().
		Only(ctx)
}

func (r *entRepository) List(ctx context.Context, limit, offset int) ([]*ent.Post, error) {
	return r.client.Post.Query().
		Limit(limit).
		Offset(offset).
		WithTags().
		Order(ent.Desc(post.FieldCreatedAt)).
		All(ctx)
}

func (r *entRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.client.Post.DeleteOneID(id).Exec(ctx)
}
