package router

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/jwt"
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
	// 1. Access Logger
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

	// 2. Official JWT Middleware
	authMiddleware, _ := jwt.New(&jwt.HertzJWTMiddleware{
		Key:         []byte(cfg.JWT.Secret),
		Timeout:     time.Duration(cfg.JWT.ExpireHours) * time.Hour,
		IdentityKey: "user_id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(string); ok {
				return jwt.MapClaims{"sub": v}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var req auth.LoginRequest
			if err := c.Bind(&req); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}
			res, err := authHandler.LoginInternal(ctx, &req)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return res.UserID, nil
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			if v, ok := data.(string); ok && v != "" {
				c.Set("user_id", v)
				return true
			}
			return false
		},
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
	})

	// 3. Auth Routes
	authGroup := h.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authMiddleware.LoginHandler)
	}

	// 4. Post Public (Custom Optimized Cache)
	posts := h.Group("/posts")
	{
		posts.GET("", middleware.CacheMiddleware(rdb, 500*time.Millisecond), postHandler.List)
		posts.GET("/:id", middleware.CacheMiddleware(rdb, 1*time.Second), postHandler.Get)
		posts.GET("/:id/comments", middleware.CacheMiddleware(rdb, 500*time.Millisecond), commentHandler.List)
	}

	// 5. Protected Actions
	protected := h.Group("")
	protected.Use(authMiddleware.MiddlewareFunc())
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
