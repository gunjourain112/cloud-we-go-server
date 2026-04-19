package post

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent"
)

type Service interface {
	CreatePost(ctx context.Context, authorID uuid.UUID, req *CreateRequest) (*Response, error)
	GetPost(ctx context.Context, id uuid.UUID) (*Response, error)
	ListPosts(ctx context.Context, limit, offset int) ([]*Response, error)
	DeletePost(ctx context.Context, id uuid.UUID, authorID uuid.UUID) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreatePost(ctx context.Context, authorID uuid.UUID, req *CreateRequest) (*Response, error) {
	p, err := s.repo.Create(ctx, req.Title, req.Body, authorID, req.Tags)
	if err != nil {
		return nil, err
	}
	return s.toResponse(p), nil
}

func (s *service) GetPost(ctx context.Context, id uuid.UUID) (*Response, error) {
	p, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.toResponse(p), nil
}

func (s *service) ListPosts(ctx context.Context, limit, offset int) ([]*Response, error) {
	posts, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	res := make([]*Response, 0, len(posts))
	for _, p := range posts {
		res = append(res, s.toResponse(p))
	}
	return res, nil
}

func (s *service) DeletePost(ctx context.Context, id uuid.UUID, authorID uuid.UUID) error {
	p, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if p.AuthorID != authorID {
		return errors.New("unauthorized")
	}

	return s.repo.Delete(ctx, id)
}

func (s *service) toResponse(p *ent.Post) *Response {
	tags := make([]string, 0, len(p.Edges.Tags))
	for _, t := range p.Edges.Tags {
		tags = append(tags, t.Name)
	}

	return &Response{
		ID:        p.ID,
		Title:     p.Title,
		Body:      p.Body,
		AuthorID:  p.AuthorID,
		Tags:      tags,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
