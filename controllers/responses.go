package controllers

import (
	"fmt"
	"indoquran-golang/forms"

	"github.com/gin-gonic/gin"
)

// DefaultResponse : default response
func DefaultResponse(c *gin.Context, status int, message string, attrib ...interface{}) {
	msg := fmt.Sprintf(message, attrib...)

	c.JSON(status, gin.H{"message": msg})
}

// TokenResponse : token result response
func TokenResponse(c *gin.Context, status int, token *forms.Token) {
	c.JSON(status, token)
}
