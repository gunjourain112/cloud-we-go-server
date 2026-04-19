package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/netpoll"
	"go.uber.org/fx"

	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/domain/comment"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/domain/post"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/infra/config"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/infra/database"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/infra/discord"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/infra/logger"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/infra/router"
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
			post.NewRepository,
			post.NewService,
			post.NewHandler,
			comment.NewRepository,
			comment.NewService,
			comment.NewHandler,
			newHertzServer,
		),
		fx.Invoke(database.RunMigration, router.RegisterRoutes, startServer),
	).Run()
}

func newHertzServer(cfg *config.Config) *server.Hertz {
	h := server.Default(
		server.WithHostPorts(fmt.Sprintf(":%d", cfg.Server.Port)),
		server.WithTransport(netpoll.NewTransporter),
	)
	return h
}

func startServer(lc fx.Lifecycle, h *server.Hertz) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go h.Run()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return h.Shutdown(ctx)
		},
	})
}
