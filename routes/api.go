package routes

import (
	"indoquran-golang/controllers"
	"indoquran-golang/middleware"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
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
		v1User.Use(middleware.TokenAuthMiddleware())
		{
			v1User.POST("/logout", user.Logout)
		}
	}
}
