package controllers

import (
	"indoquran-golang/forms"
	"indoquran-golang/models"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
)

// Import the userModel from the models
var userModel = new(models.UserModel)

// UserController defines the user controller methods
type UserController struct{}

// Signup controller handles registering a user
func (u *UserController) Signup(c *gin.Context) {
	var data forms.SignupUserCommand

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()
		return
	}

	// email validator
	emailErr := checkmail.ValidateFormat(data.Email)
	if emailErr != nil {
		c.JSON(400, gin.H{"message": "Email is invalid " + data.Email})
		c.Abort()
		return
	}

	err := userModel.Signup(data)

	// Check if there was an error when saving user
	if err != nil {
		c.JSON(400, gin.H{"message": "Problem creating an account"})
		c.Abort()
		return
	}

	c.JSON(201, gin.H{"message": "New user account registered"})
}
