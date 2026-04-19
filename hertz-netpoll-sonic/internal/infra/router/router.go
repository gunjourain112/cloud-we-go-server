package router

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/domain/auth"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/domain/comment"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/domain/post"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/infra/config"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/infra/middleware"
	"github.com/redis/go-redis/v9"
)

func RegisterRoutes(
	h *server.Hertz,
	rdb *redis.Client,
	cfg *config.Config,
	authHandler *auth.Handler,
	postHandler *post.Handler,
	commentHandler *comment.Handler,
) {
	authGroup := h.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	posts := h.Group("/posts")
	{
		posts.GET("", middleware.CacheMiddleware(rdb, 500*time.Millisecond), postHandler.List)
		posts.GET("/:id", middleware.CacheMiddleware(rdb, 1*time.Second), postHandler.Get)
		posts.GET("/:id/comments", middleware.CacheMiddleware(rdb, 500*time.Millisecond), commentHandler.List)
	}

	protected := h.Group("", middleware.AuthMiddleware(cfg))
	{
		protected.POST("/posts", postHandler.Create)
		protected.DELETE("/posts/:id", postHandler.Delete)
		protected.POST("/posts/:id/comments", commentHandler.Create)
		protected.POST("/posts/:id/comments/:cid/replies", commentHandler.Reply)
		protected.DELETE("/posts/:id/comments/:cid", commentHandler.Delete)
	}

	h.GET("/health", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(200, map[string]string{"status": "ok"})
	})
}
