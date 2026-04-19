package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/redis/go-redis/v9"
)

func RateLimitMiddleware(rdb *redis.Client, limit int) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var key string
		if uid, exists := c.Get("user_id"); exists {
			key = fmt.Sprintf("limit:%v", uid)
		} else {
			key = fmt.Sprintf("limit:%s", c.ClientIP())
		}

		count, _ := rdb.Incr(ctx, key).Result()
		if count == 1 {
			rdb.Expire(ctx, key, time.Second)
		}

		if count > int64(limit) {
			c.AbortWithStatus(429)
			return
		}

		c.Next(ctx)
	}
}
