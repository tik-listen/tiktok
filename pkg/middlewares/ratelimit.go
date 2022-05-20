package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// If it can't take token then return res
		// If bucket.Take(1) > 0 {
		if bucket.TakeAvailable(1) != 1 {
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		// take next
		c.Next()
	}
}
