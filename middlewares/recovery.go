package middlewares

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recv := recover(); recv != nil {
				if err, ok := recv.(error); ok {
					_ = c.Error(err)
				} else if err, ok := recv.(string); ok {
					_ = c.Error(errors.New(err))
				} else {
					_ = c.Error(fmt.Errorf("unknown error: %v", recv))
				}
			}
		}()
		c.Next()
	}
}
