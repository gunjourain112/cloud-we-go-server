package router

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/zsais/go-gin-prometheus"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"

	"github.com/gunjourain112/cloud-we-go-server/gin/internal/domain/auth"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/domain/comment"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/domain/post"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/infra/config"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/infra/middleware"
)

func NewEngine(cfg *config.Config) *gin.Engine {
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	
	return r
}

func RegisterRoutes(
	r *gin.Engine,
	log *zap.Logger,
	db *sql.DB,
	rdb *redis.Client,
	mdb *mongo.Database,
	cfg *config.Config,
	authHandler *auth.Handler,
	postHandler *post.Handler,
	commentHandler *comment.Handler,
) {
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		log.Info("Request",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Duration("latency", latency),
		)
	})

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	posts := r.Group("/posts")
	{
		posts.GET("", middleware.CacheMiddleware(rdb, 500*time.Millisecond), postHandler.List)
		posts.GET("/:id", middleware.CacheMiddleware(rdb, 1*time.Second), postHandler.Get)
		posts.GET("/:id/comments", middleware.CacheMiddleware(rdb, 500*time.Millisecond), commentHandler.List)
	}

	// Remove AuthMiddleware
	publicActions := r.Group("")
	{
		publicActions.POST("/posts", postHandler.Create)
		publicActions.DELETE("/posts/:id", postHandler.Delete)
		publicActions.POST("/posts/:id/comments", commentHandler.Create)
		publicActions.POST("/posts/:id/comments/:cid/replies", commentHandler.Reply)
		publicActions.DELETE("/posts/:id/comments/:cid", commentHandler.Delete)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}
