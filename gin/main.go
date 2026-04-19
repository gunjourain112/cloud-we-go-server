package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/fx"
	"go.uber.org/zap"
)
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/config"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/database"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/logger"
)

func main() {
	fx.New(
		fx.Provide(
			config.Load,
			logger.NewLogger,
			database.NewPostgres,
			database.NewRedis,
			database.NewMongo,
			database.NewEntClient,
			newGinEngine,
		),
		fx.Invoke(registerRoutes, startServer),
	).Run()
}

func newGinEngine(cfg *config.Config, log *zap.Logger) *gin.Engine {
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	// TODO: Add custom zap logger middleware and other common middlewares here
	
	return r
}

func registerRoutes(
	r *gin.Engine,
	log *zap.Logger,
	db *sql.DB,
	rdb *redis.Client,
	mdb *mongo.Database,
) {
	r.GET("/health", func(c *gin.Context) {
		status := gin.H{
			"status": "ok",
			"db":     "ok",
			"redis":  "ok",
			"mongo":  "ok",
		}

		if err := db.Ping(); err != nil {
			status["db"] = "error: " + err.Error()
		}
		if err := rdb.Ping(c).Err(); err != nil {
			status["redis"] = "error: " + err.Error()
		}
		if err := mdb.Client().Ping(c, nil); err != nil {
			status["mongo"] = "error: " + err.Error()
		}

		c.JSON(http.StatusOK, status)
	})
}

func startServer(lc fx.Lifecycle, r *gin.Engine, cfg *config.Config, log *zap.Logger) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting Gin server", zap.Int("port", cfg.Server.Port))
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatal("Failed to serve", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Shutting down Gin server")
			return srv.Shutdown(ctx)
		},
	})
}
