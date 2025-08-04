package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireJSONContentType() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost ||
			c.Request.Method == http.MethodPut {
			if c.ContentType() != "application/json" {
				c.AbortWithStatusJSON(http.StatusUnsupportedMediaType, gin.H{
					"error": "Content-Type must be application/json",
				})
				return
			}
			c.Next()
		}

	}
}
