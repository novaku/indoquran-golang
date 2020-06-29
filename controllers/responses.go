package controllers

import (
	"fmt"
	"indoquran-golang/models"

	"github.com/gin-gonic/gin"
)

// DefaultResponse : default response
func DefaultResponse(c *gin.Context, status int, data interface{}, message string, attrib ...interface{}) {
	msg := fmt.Sprintf(message, attrib...)

	c.JSON(status, gin.H{
		"data":    data,
		"message": msg,
	})
}

// TokenResponse : token result response
func TokenResponse(c *gin.Context, status int, token *models.Token) {
	c.JSON(status, token)
}
