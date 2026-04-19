package database

import (
	"database/sql"

	"context"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/ent"
	"github.com/gunjourain112/cloud-we-go-server/hertz/internal/infra/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewEntClient(db *sql.DB) *ent.Client {
	driver := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(driver))
}

func RunMigration(lc fx.Lifecycle, client *ent.Client, cfg *config.Config, log *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if cfg.App.Env == "development" && cfg.Postgres.AutoMigrate {
				log.Info("Running auto migration")
				if err := client.Schema.Create(ctx); err != nil {
					return fmt.Errorf("failed creating schema resources: %w", err)
				}
			}
			return nil
		},
	})
}
