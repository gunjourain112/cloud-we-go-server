package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	App       AppConfig
	Server    ServerConfig
	Postgres  PostgresConfig
	Redis     RedisConfig
	Mongo     MongoConfig
	JWT       JWTConfig
	RateLimit RateLimitConfig
	Discord   DiscordConfig
}

type AppConfig struct {
	Env string `env:"APP_ENV" envDefault:"development"`
}

type ServerConfig struct {
	Port int `env:"GIN_PORT" envDefault:"8080"`
}

type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port     int    `env:"POSTGRES_PORT" envDefault:"5432"`
	User     string `env:"POSTGRES_USER" envDefault:"cwg"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"cwg"`
	DB       string `env:"POSTGRES_DB"       envDefault:"cwg"`
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST"     envDefault:"localhost"`
	Port     int    `env:"REDIS_PORT"     envDefault:"6379"`
	Password string `env:"REDIS_PASSWORD" envDefault:""`
	DB       int    `env:"REDIS_DB"       envDefault:"0"`
}

type MongoConfig struct {
	URI string `env:"MONGO_URI" envDefault:"mongodb://cwg:cwg@localhost:27017"`
	DB  string `env:"MONGO_DB"  envDefault:"cwg"`
}

type JWTConfig struct {
	Secret      string `env:"JWT_SECRET"       envDefault:"change_me"`
	ExpireHours int    `env:"JWT_EXPIRE_HOURS" envDefault:"24"`
}

type RateLimitConfig struct {
	RPS   int `env:"RATE_LIMIT_RPS"   envDefault:"100"`
	Burst int `env:"RATE_LIMIT_BURST" envDefault:"200"`
}

type DiscordConfig struct {
	WebhookURL string `env:"DISCORD_WEBHOOK_URL" envDefault:""`
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
