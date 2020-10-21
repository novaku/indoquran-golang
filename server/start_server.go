package server

import (
	"bitbucket.org/indoquran-api/config"
	"bitbucket.org/indoquran-api/routes"

	"github.com/gin-gonic/gin"
)

// StartServer : start the server, load the router
func StartServer() {
	r := gin.Default()
	// gin.SetMode(gin.DebugMode)
	gin.SetMode(config.LoadConfig().Server.SetMode)

	routes.LoadRoutes(r)

	// Handle error response when a route is not defined
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	r.Run(config.LoadConfig().Server.Host + ":" + config.LoadConfig().Server.Port)
}
