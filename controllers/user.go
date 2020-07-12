package controllers

import (
	"indoquran-golang/config"
	"indoquran-golang/helpers"
	"indoquran-golang/models"
	"indoquran-golang/models/modelstruct"
	"net/http"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/golang/glog"
)

// Import the userModel from the models
var userModel = new(models.UserModel)
var client *redis.Client

// UserController defines the user controller methods
type UserController struct{}

func init() {
	//Initializing redis
	dsn := config.LoadConfig().Redis.Host + ":" + config.LoadConfig().Redis.Port
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

// Signup controller handles registering a user
func (u *UserController) Signup(c *gin.Context) {
	var data modelstruct.SignupUserCommand

	if c.BindJSON(&data) != nil {
		DefaultResponse(c, http.StatusNotAcceptable, data, "Provide relevant fields")
		c.Abort()
		return
	}

	// email validator
	emailErr := checkmail.ValidateFormat(data.Email)
	if emailErr != nil {
		DefaultResponse(c, http.StatusBadRequest, data, "Email is invalid (%s)", data.Email)
		c.Abort()
		return
	}

	// search if email already registered
	resEmail, _ := userModel.GetUserByEmail(data.Email)
	if resEmail.Email != "" {
		DefaultResponse(c, http.StatusForbidden, data, "Email %s already in use!", data.Email)
		c.Abort()
		return
	}

	err := userModel.Signup(data)

	// Check if there was an error when saving user
	if err != nil {
		DefaultResponse(c, http.StatusBadRequest, data, "Problem creating an account")
		c.Abort()
		return
	}

	DefaultResponse(c, http.StatusCreated, data, "New account registered (%s)", data.Email)
}

// Login allows a user to login a user and get
// access token
func (u *UserController) Login(c *gin.Context) {
	var data modelstruct.LoginUserCommand
	var tokenResult = &modelstruct.TokenResult{}

	// Bind the request body data to var data and check if all details are provided
	if c.BindJSON(&data) != nil {
		DefaultResponse(c, http.StatusNotAcceptable, data, "Request provided is not acceptable")
		c.Abort()
		return
	}

	user, errEmail := userModel.GetUserByEmail(data.Email)

	if user.Email == "" {
		DefaultResponse(c, http.StatusNotFound, data, "User email %s account not found", data.Email)
		c.Abort()
		return
	}

	if errEmail != nil {
		DefaultResponse(c, http.StatusBadRequest, data, "Problem logging into your account")
		c.Abort()
		return
	}

	// Get the hashed password from the saved document
	hashedPassword := []byte(user.Password)
	// Get the password provided in the request.body
	password := []byte(data.Password)

	errHash := helpers.PasswordCompare(hashedPassword, password)

	if errHash != nil {
		DefaultResponse(c, http.StatusForbidden, data, "Invalid user credentials")
		c.Abort()
		return
	}

	ts, err := helpers.CreateToken(user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := CreateAuth(user.ID.Hex(), ts)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}

	tokenResult.AccessToken = ts.AccessToken
	tokenResult.RefreshToken = ts.RefreshToken

	TokenResponse(c, http.StatusOK, tokenResult)
}

// CreateAuth : set to redis for login credential
func CreateAuth(userID string, td *modelstruct.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := client.Set(td.AccessUUID, userID, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := client.Set(td.RefreshUUID, userID, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

// Logout : logout the user
func (u *UserController) Logout(c *gin.Context) {
	au, err := helpers.ExtractTokenMetadata(c.Request)
	if err != nil {
		glog.Errorf("%s", c.Request.Header.Get("Authorization"))
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	deleted, delErr := helpers.DeleteAuth(au.AccessUUID)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}
