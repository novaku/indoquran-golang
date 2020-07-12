package routes

import (
	"indoquran-golang/controllers"
	"indoquran-golang/middleware"

	"github.com/gin-gonic/gin"
)

// APIRoutes : routes for API
func APIRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		// No auth
		user := new(controllers.UserController)
		v1User := v1.Group("user")
		{
			v1User.POST("/signup", user.Signup)
			v1User.POST("/login", user.Login)
		}

		// With auth
		v1Auth := v1.Use(middleware.TokenAuthMiddleware())
		v1Auth.POST("/user/logout", user.Logout)

	}
}
