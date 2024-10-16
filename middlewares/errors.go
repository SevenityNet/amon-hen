package middlewares

import (
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"runtime"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			log.Warnf("%d error occurred in request:", len(c.Errors))
			for i, err := range c.Errors {
				log.Errorf("Error %d: %s\n%s", i, err.Error(), getStack())
			}

			c.JSON(500, gin.H{
				"error": "Internal Server Error",
			})

			c.Abort()
		}
	}
}

func getStack() string {
	stack := make([]byte, 1024)
	stack = stack[:runtime.Stack(stack, false)]
	return string(stack)
}
