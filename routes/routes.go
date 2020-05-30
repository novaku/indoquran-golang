package routes

import "github.com/gin-gonic/gin"

// LoadRoutes : load the routing list
func LoadRoutes(r *gin.Engine) {
	APIRoutes(r)
}
