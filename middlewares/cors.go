package middlewares

import (
	"github.com/gin-gonic/gin"
	"os"
)

func CORS() gin.HandlerFunc {
	corsAllowedOrigin := os.Getenv("CORS_ALLOWED_ORIGINS")

	return func(c *gin.Context) {
		if corsAllowedOrigin == "*" {
			requestDomain := c.Request.Header.Get("Origin")
			c.Writer.Header().Set("Access-Control-Allow-Origin", requestDomain)
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", corsAllowedOrigin)
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PATCH, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Baggage, Accept, Sentry-Trace")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
