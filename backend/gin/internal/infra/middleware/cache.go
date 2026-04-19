package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func CacheMiddleware(rdb *redis.Client, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := fmt.Sprintf("cache:%s", c.Request.RequestURI)

		if val := rdb.Get(c, key).Val(); val != "" {
			c.Header("X-Cache", "HIT")
			c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(val))
			c.Abort()
			return
		}

		w := &responseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = w

		c.Next()

		if c.Request.Method == http.MethodGet && c.Writer.Status() == http.StatusOK {
			rdb.Set(c, key, w.body.String(), duration)
		}
	}
}
