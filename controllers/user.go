package controllers

import (
	"indoquran-golang/helpers"
	"indoquran-golang/models"
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
)

// Import the userModel from the models
var userModel = new(models.UserModel)

// UserController defines the user controller methods
type UserController struct{}

// Signup controller handles registering a user
func (u *UserController) Signup(c *gin.Context) {
	var data models.SignupUserCommand

	if c.BindJSON(&data) != nil {
		DefaultResponse(c, http.StatusNotAcceptable, "Provide relevant fields")
		c.Abort()
		return
	}

	// email validator
	emailErr := checkmail.ValidateFormat(data.Email)
	if emailErr != nil {
		DefaultResponse(c, http.StatusBadRequest, "Email is invalid (%s)", data.Email)
		c.Abort()
		return
	}

	// search if email already registered
	resEmail, _ := userModel.GetUserByEmail(data.Email)
	if resEmail.Email != "" {
		DefaultResponse(c, http.StatusForbidden, "Email %s already in use!", data.Email)
		c.Abort()
		return
	}

	err := userModel.Signup(data)

	// Check if there was an error when saving user
	if err != nil {
		DefaultResponse(c, http.StatusBadRequest, "Problem creating an account")
		c.Abort()
		return
	}

	DefaultResponse(c, http.StatusCreated, "New account registered (%s)", data.Email)
}

// Login allows a user to login a user and get
// access token
func (u *UserController) Login(c *gin.Context) {
	var data models.LoginUserCommand

	// Bind the request body data to var data and check if all details are provided
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	result, err := userModel.GetUserByEmail(data.Email)

	if result.Email == "" {
		DefaultResponse(c, http.StatusNotFound, "User email %s account was not found", data.Email)
		c.Abort()
		return
	}

	if err != nil {
		DefaultResponse(c, http.StatusBadRequest, "Problem logging into your account")
		c.Abort()
		return
	}

	// Get the hashed password from the saved document
	hashedPassword := []byte(result.Password)
	// Get the password provided in the request.body
	password := []byte(data.Password)

	err = helpers.PasswordCompare(password, hashedPassword)

	if err != nil {
		DefaultResponse(c, http.StatusForbidden, "Invalid user credentials")
		c.Abort()
		return
	}

	jwtToken, err2 := helpers.GenerateToken(data.Email)

	// If we fail to generate token for access
	if err2 != nil {
		DefaultResponse(c, http.StatusInternalServerError, "There was a problem logging you in, try again later")
		c.Abort()
		return
	}

	TokenResponse(c, http.StatusOK, &models.Token{
		Message: "Log in success",
		Token:   jwtToken,
	})
}
