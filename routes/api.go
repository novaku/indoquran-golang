package routes

import (
	"net/http"

	"bitbucket.org/indoquran-api/controllers"
	"bitbucket.org/indoquran-api/helpers"
	"bitbucket.org/indoquran-api/helpers/logger"
	"bitbucket.org/indoquran-api/middleware"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

const (
	infoLogTag = "routes|api.go|APIRoutes()"
)

// APIRoutes : routes for API
func APIRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.Use(requestid.New())
	{
		// No auth
		user := new(controllers.UserController)
		v1User := v1.Group("user")
		{
			v1User.POST("/signup", user.Signup)
			v1User.POST("/login", user.Login)
		}

		// With auth
		v1.Use(middleware.TokenAuthMiddleware())
		{
			v1User.GET("/logout", user.Logout)
			v1.GET("/info", func(c *gin.Context) {
				requestID := requestid.Get(c)
				au, err := helpers.ExtractTokenMetadata(c.Request, requestID)
				if err != nil {
					logger.Error(infoLogTag, requestID, "Unable to extract token metadata on logout system, error: %+v", err)
					c.JSON(http.StatusUnauthorized, "unauthorized")
					return
				}

				c.JSON(200, gin.H{
					"access_id": au.AccessUUID,
					"user_id":   au.UserID,
				})
			})
		}
	}
}
