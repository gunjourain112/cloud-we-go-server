package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/gunjourain112/cloud-we-go-server/gin/internal/domain/auth"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/domain/comment"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/domain/post"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/domain/user"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/infra/config"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/infra/database"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/infra/discord"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/infra/logger"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/infra/router"
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
			discord.NewClient,
			user.NewRepository,
			auth.NewService,
			auth.NewHandler,
			post.NewRepository,
			post.NewService,
			post.NewHandler,
			comment.NewRepository,
			comment.NewService,
			comment.NewHandler,
			router.NewEngine,
		),
		fx.Invoke(database.RunMigration, router.RegisterRoutes, startServer),
	).Run()
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
