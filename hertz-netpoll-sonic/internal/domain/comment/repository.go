package comment

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CommentDoc struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	PostID    uuid.UUID     `bson:"post_id"`
	AuthorID  uuid.UUID     `bson:"author_id"`
	Body      string        `bson:"body"`
	Replies   []ReplyDoc    `bson:"replies"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}

type ReplyDoc struct {
	ID        bson.ObjectID `bson:"_id"`
	AuthorID  uuid.UUID     `bson:"author_id"`
	Body      string        `bson:"body"`
	CreatedAt time.Time     `bson:"created_at"`
}

type Repository interface {
	Create(ctx context.Context, doc *CommentDoc) error
	GetByPostID(ctx context.Context, postID uuid.UUID) ([]*CommentDoc, error)
	AddReply(ctx context.Context, commentID bson.ObjectID, reply *ReplyDoc) error
	Delete(ctx context.Context, id bson.ObjectID, authorID uuid.UUID) error
}

type mongoRepository struct {
	col *mongo.Collection
}

func NewRepository(db *mongo.Database) Repository {
	return &mongoRepository{col: db.Collection("comments")}
}

func (r *mongoRepository) Create(ctx context.Context, doc *CommentDoc) error {
	doc.CreatedAt = time.Now()
	doc.UpdatedAt = time.Now()
	_, err := r.col.InsertOne(ctx, doc)
	return err
}

func (r *mongoRepository) GetByPostID(ctx context.Context, postID uuid.UUID) ([]*CommentDoc, error) {
	cursor, err := r.col.Find(ctx, bson.M{"post_id": postID})
	if err != nil {
		return nil, err
	}
	var docs []*CommentDoc
	err = cursor.All(ctx, &docs)
	return docs, err
}

func (r *mongoRepository) AddReply(ctx context.Context, commentID bson.ObjectID, reply *ReplyDoc) error {
	reply.ID = bson.NewObjectID()
	reply.CreatedAt = time.Now()
	_, err := r.col.UpdateOne(ctx, bson.M{"_id": commentID}, bson.M{"$push": bson.M{"replies": reply}})
	return err
}

func (r *mongoRepository) Delete(ctx context.Context, id bson.ObjectID, authorID uuid.UUID) error {
	_, err := r.col.DeleteOne(ctx, bson.M{"_id": id, "author_id": authorID})
	return err
}
