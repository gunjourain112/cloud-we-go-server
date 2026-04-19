package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"

	"github.com/gunjourain112/cloud-we-go-server/gin/internal/infra/config"
	"github.com/redis/go-redis/v9"
	"time"
)

func RateLimitMiddleware(cfg *config.Config, rdb *redis.Client, limit int) gin.HandlerFunc {
	rate := limiter.Rate{
		Limit:  int64(limit),
		Period: time.Second,
	}

	store, err := sredis.NewStoreWithOptions(rdb, limiter.StoreOptions{
		Prefix: "rate-limit:",
	})
	if err != nil {
		panic(err)
	}

	middleware := mgin.NewMiddleware(limiter.New(store, rate), mgin.WithKeyGetter(func(c *gin.Context) string {
		if uid, exists := c.Get("user_id"); exists {
			return uid.(string)
		}
		return c.ClientIP()
	}))

	return middleware
}
