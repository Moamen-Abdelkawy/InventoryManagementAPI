package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// limiter enforces 1 request per second with a burst capacity of 5.
var limiter = rate.NewLimiter(1, 5)

// RateLimitMiddleware restricts incoming HTTP requests according to the token bucket rule.
func RateLimitMiddleware(c *gin.Context) {
	if !limiter.Allow() {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"error": "Too many requests. Please try again later.",
		})
		return
	}
	c.Next()
}
