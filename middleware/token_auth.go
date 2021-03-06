package middleware

import (
	"net/http"

	"bitbucket.org/indoquran-api/helpers"

	"github.com/gin-gonic/gin"
)

// TokenAuthMiddleware : middleware for token auth checkk
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helpers.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
