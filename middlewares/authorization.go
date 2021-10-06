package middlewares

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		userData := c.MustGet("userData").(jwt.MapClaims)
		userRole := userData["role"].(string)

		if userRole != "admin" {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Unauthorized",
				"message": "you are not authorized to access this data",
			})
			return
		}
		c.Next()
	}
}
