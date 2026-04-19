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
	"go.uber.org/zap"
)

func RegisterRoutes(
	h *server.Hertz,
	rdb *redis.Client,
	cfg *config.Config,
	log *zap.Logger,
	authHandler *auth.Handler,
	postHandler *post.Handler,
	commentHandler *comment.Handler,
) {
	h.Use(func(ctx context.Context, c *app.RequestContext) {
		start := time.Now()
		c.Next(ctx)
		latency := time.Since(start)
		log.Info("Request",
			zap.Int("status", c.Response.StatusCode()),
			zap.String("method", string(c.Request.Header.Method())),
			zap.String("path", string(c.Request.Path())),
			zap.Duration("latency", latency),
		)
	})

	authGroup := h.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.LoginInternal) // Using internal for direct response
	}

	posts := h.Group("/posts")
	{
		posts.GET("", middleware.CacheMiddleware(rdb, 500*time.Millisecond), postHandler.List)
		posts.GET("/:id", middleware.CacheMiddleware(rdb, 1*time.Second), postHandler.Get)
		posts.GET("/:id/comments", middleware.CacheMiddleware(rdb, 500*time.Millisecond), commentHandler.List)
	}

	// Remove AuthMiddleware
	publicActions := h.Group("")
	{
		publicActions.POST("/posts", postHandler.Create)
		publicActions.DELETE("/posts/:id", postHandler.Delete)
		publicActions.POST("/posts/:id/comments", commentHandler.Create)
		publicActions.POST("/posts/:id/comments/:cid/replies", commentHandler.Reply)
		publicActions.DELETE("/posts/:id/comments/:cid", commentHandler.Delete)
	}

	h.GET("/health", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(200, map[string]string{"status": "ok"})
	})
}
