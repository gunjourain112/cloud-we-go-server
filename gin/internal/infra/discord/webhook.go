package discord

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/infra/config"
)

type Client struct {
	cfg *config.Config
	rdb *redis.Client
}

func NewClient(cfg *config.Config, rdb *redis.Client) *Client {
	return &Client{cfg: cfg, rdb: rdb}
}

func (c *Client) SendError(ctx context.Context, message string) error {
	if c.cfg.Discord.WebhookURL == "" {
		return nil
	}

	key := "discord:cooldown:error"
	if c.rdb.Get(ctx, key).Val() != "" {
		return nil
	}

	payload := map[string]string{"content": fmt.Sprintf("🚨 **[Gin Error]** %s", message)}
	body, _ := json.Marshal(payload)

	resp, err := http.Post(c.cfg.Discord.WebhookURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	c.rdb.Set(ctx, key, "sent", 60*time.Second)
	return nil
}
