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
	// Public Group
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

	// Protected Group
	protected := r.Group("")
	protected.Use(middleware.AuthMiddleware(cfg))
	{
		protected.POST("/posts", postHandler.Create)
		protected.DELETE("/posts/:id", postHandler.Delete)
		protected.POST("/posts/:id/comments", commentHandler.Create)
		protected.POST("/posts/:id/comments/:cid/replies", commentHandler.Reply)
		protected.DELETE("/posts/:id/comments/:cid", commentHandler.Delete)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}
