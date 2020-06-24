package server

import (
	"indoquran-golang/config"
	"indoquran-golang/routes"
	"time"

	"github.com/gin-gonic/gin"
	ginglog "github.com/szuecs/gin-glog"
)

// StartServer : start the server, load the router
func StartServer() {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.Use(ginglog.Logger(2 * time.Second))

	routes.LoadRoutes(r)

	// Handle error response when a route is not defined
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	r.Run(config.LoadConfig().Server.Host + ":" + config.LoadConfig().Server.Port)
}
