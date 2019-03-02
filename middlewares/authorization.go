package middlewares

import (
	"github.com/gin-gonic/gin"
)

func VerifyJwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
