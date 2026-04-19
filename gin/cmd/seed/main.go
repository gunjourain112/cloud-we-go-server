package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"

	"github.com/gunjourain112/cloud-we-go-server/gin/internal/domain/comment"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/domain/post"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/domain/user"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/infra/config"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/infra/database"
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

	userRepo := user.NewRepository(entClient)
	postRepo := post.NewRepository(entClient)
	commentRepo := comment.NewRepository(mdb)

	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	pwHash := string(password)

	fmt.Println("🌱 Seeding data...")

	for i := 1; i <= 10; i++ {
		u, err := userRepo.Create(ctx, fmt.Sprintf("user%d@example.com", i), pwHash, fmt.Sprintf("User%d", i))
		if err != nil {
			fmt.Printf("User %d already exists, skipping\n", i)
			u, _ = userRepo.GetByEmail(ctx, fmt.Sprintf("user%d@example.com", i))
		}

		for j := 1; j <= 5; j++ {
			p, err := postRepo.Create(ctx,
				fmt.Sprintf("Post %d by User %d", j, i),
				"This is a sample post body content for benchmarking purposes. "+time.Now().String(),
				u.ID,
				[]string{"golang", "gin", "benchmark", fmt.Sprintf("tag%d", j)},
			)
			if err != nil {
				log.Fatal(err)
			}

			for k := 1; k <= 3; k++ {
				c := &comment.CommentDoc{
					PostID:   p.ID,
					AuthorID: u.ID,
					Body:     fmt.Sprintf("Comment %d on post %d", k, j),
					Replies: []comment.ReplyDoc{
						{
							ID:        bson.NewObjectID(),
							AuthorID:  u.ID,
							Body:      "Sample reply content",
							CreatedAt: time.Now(),
						},
					},
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
				commentRepo.Create(ctx, c)
			}
		}
	}

	fmt.Println("✅ Seeding completed!")
}
