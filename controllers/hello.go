package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloWorldController will hold the methods to the
type HelloWorldController struct{}

// Default controller handles returning the hello world JSON response
func (h *HelloWorldController) Default(c *gin.Context) {
	DefaultResponse(c, http.StatusOK, "Hello %s", "world")
}
