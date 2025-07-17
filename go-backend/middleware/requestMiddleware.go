package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		timestamp := c.GetHeader("X-Timestamp")
		if timestamp == "" {
			timestamp = time.Now().Format(time.RFC3339)
		}

		combined := requestID + "-" + timestamp

		c.Set("request_id", combined)

		c.Writer.Header().Set("Request-ID", combined)
		c.Request.Header.Set("Request-ID", combined)

		c.Next()
	}
}
