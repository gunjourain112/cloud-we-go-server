package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/gunjourain112/cloud-we-go-server/gin/internal/domain/comment"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/domain/post"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/infra/config"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/infra/database"
	"github.com/google/uuid"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	cfg.Postgres.AutoMigrate = true

	ctx := context.Background()
	db, _ := database.NewPostgres(cfg)
	entClient := database.NewEntClient(db)
	mdb, _ := database.NewMongo(cfg)

	if err := entClient.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema: %v", err)
	}

	postRepo := post.NewRepository(entClient)
	commentRepo := comment.NewRepository(mdb)

	testUserID := uuid.MustParse("00000000-0000-0000-0000-000000000001")

	fmt.Println("🌱 Seeding data...")

	for i := 1; i <= 50; i++ {
		p, err := postRepo.Create(ctx, 
			fmt.Sprintf("Raw Post %d", i),
			"Benchmark content without user joins.",
			testUserID,
			[]string{"golang", "raw", "benchmark"},
		)
		if err != nil {
			log.Fatal(err)
		}

		for k := 1; k <= 3; k++ {
			c := &comment.CommentDoc{
				PostID:   p.ID,
				AuthorID: testUserID,
				Body:     fmt.Sprintf("Comment %d on post %d", k, i),
				Replies:  []comment.ReplyDoc{},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			commentRepo.Create(ctx, c)
		}
	}

	fmt.Println("✅ Seeding completed!")
}
