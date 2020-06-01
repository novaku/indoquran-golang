package routes

import (
	"indoquran-golang/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIRoutes routes for API
func APIRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong kaleee"})
		})

		user := new(controllers.UserController)
		v1.POST("/signup", user.Signup)
		v1.POST("/login", user.Login)
	}
}
