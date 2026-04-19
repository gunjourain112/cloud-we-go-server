package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad_Defaults(t *testing.T) {
	cfg, err := Load()
	require.NoError(t, err)

	assert.Equal(t, "development", cfg.App.Env)
	assert.Equal(t, 8080, cfg.Server.Port)
	assert.Equal(t, "localhost", cfg.Postgres.Host)
	assert.Equal(t, 5432, cfg.Postgres.Port)
	assert.Equal(t, "localhost", cfg.Redis.Host)
	assert.Equal(t, 6379, cfg.Redis.Port)
	assert.Equal(t, 24, cfg.JWT.ExpireHours)
	assert.Equal(t, 100, cfg.RateLimit.RPS)
	assert.Equal(t, 200, cfg.RateLimit.Burst)
}

func TestLoad_EnvOverride(t *testing.T) {
	t.Setenv("APP_ENV", "production")
	t.Setenv("GIN_PORT", "9090")
	t.Setenv("POSTGRES_HOST", "db.example.com")
	t.Setenv("JWT_SECRET", "supersecret")
	t.Setenv("JWT_EXPIRE_HOURS", "48")

	cfg, err := Load()
	require.NoError(t, err)

	assert.Equal(t, "production", cfg.App.Env)
	assert.Equal(t, 9090, cfg.Server.Port)
	assert.Equal(t, "db.example.com", cfg.Postgres.Host)
	assert.Equal(t, "supersecret", cfg.JWT.Secret)
	assert.Equal(t, 48, cfg.JWT.ExpireHours)
}
