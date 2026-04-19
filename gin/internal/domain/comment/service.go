package comment

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Service interface {
	CreateComment(ctx context.Context, postID, authorID uuid.UUID, req *CreateRequest) (*Response, error)
	GetCommentsByPost(ctx context.Context, postID uuid.UUID) ([]*Response, error)
	AddReply(ctx context.Context, commentID bson.ObjectID, authorID uuid.UUID, req *ReplyRequest) error
	DeleteComment(ctx context.Context, id bson.ObjectID, authorID uuid.UUID) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateComment(ctx context.Context, postID, authorID uuid.UUID, req *CreateRequest) (*Response, error) {
	doc := &CommentDoc{
		PostID:   postID,
		AuthorID: authorID,
		Body:     req.Body,
		Replies:  []ReplyDoc{},
	}
	if err := s.repo.Create(ctx, doc); err != nil {
		return nil, err
	}
	return s.toResponse(doc), nil
}

func (s *service) GetCommentsByPost(ctx context.Context, postID uuid.UUID) ([]*Response, error) {
	docs, err := s.repo.GetByPostID(ctx, postID)
	if err != nil {
		return nil, err
	}
	res := make([]*Response, 0, len(docs))
	for _, doc := range docs {
		res = append(res, s.toResponse(doc))
	}
	return res, nil
}

func (s *service) AddReply(ctx context.Context, commentID bson.ObjectID, authorID uuid.UUID, req *ReplyRequest) error {
	reply := &ReplyDoc{
		AuthorID: authorID,
		Body:     req.Body,
	}
	return s.repo.AddReply(ctx, commentID, reply)
}

func (s *service) DeleteComment(ctx context.Context, id bson.ObjectID, authorID uuid.UUID) error {
	return s.repo.Delete(ctx, id, authorID)
}

func (s *service) toResponse(doc *CommentDoc) *Response {
	replies := make([]Reply, 0, len(doc.Replies))
	for _, r := range doc.Replies {
		replies = append(replies, Reply{
			ID:        r.ID.Hex(),
			AuthorID:  r.AuthorID.String(),
			Body:      r.Body,
			CreatedAt: r.CreatedAt,
		})
	}

	return &Response{
		ID:        doc.ID.Hex(),
		PostID:    doc.PostID.String(),
		AuthorID:  doc.AuthorID.String(),
		Body:      doc.Body,
		Replies:   replies,
		CreatedAt: doc.CreatedAt,
		UpdatedAt: doc.UpdatedAt,
	}
}
