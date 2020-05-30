package routes

import (
	"indoquran-golang/controllers"

	"github.com/gin-gonic/gin"
)

// APIRoutes routes for API
func APIRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		// Define the hello controller
		hello := new(controllers.HelloWorldController)
		// Define a GET request to call the Default
		// method in controllers/hello.go
		v1.GET("/hello", hello.Default)

		user := new(controllers.UserController)
		v1.POST("/signup", user.Signup)
	}
}
