package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/redis/go-redis/v9"
)

func CacheMiddleware(rdb *redis.Client, duration time.Duration) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		key := fmt.Sprintf("cache:%s", string(c.Request.RequestURI()))

		if val, err := rdb.Get(ctx, key).Result(); err == nil && val != "" {
			c.Header("X-Cache", "HIT")
			c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(val))
			c.Abort()
			return
		}

		c.Next(ctx)

		if c.Response.StatusCode() == http.StatusOK {
			rdb.Set(ctx, key, c.Response.Body(), duration)
		}
	}
}
