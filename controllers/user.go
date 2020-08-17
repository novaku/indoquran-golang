package controllers

import (
	"indoquran-golang/config"
	"indoquran-golang/helpers"
	"indoquran-golang/helpers/logger"
	"indoquran-golang/models"
	"indoquran-golang/models/modelstruct"
	"net/http"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
)

// Import the userModel from the models
var userModel = new(models.UserModel)
var client *redis.Client

const (
	signupLogTag = "controllers|user.go|Signup()"
	loginLogTag  = "controllers|user.go|Login()"
	logoutLogTag = "controllers|user.go|Logout"
)

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
	requestID := requestid.Get(c)
	var data modelstruct.SignupUserCommand

	if bindErr := c.BindJSON(&data); bindErr != nil {
		logger.Error(signupLogTag, requestID, "signup JSON bind failed, error : %+v", bindErr)
		DefaultResponse(c, http.StatusNotAcceptable, data, "Provide relevant fields")
		c.Abort()
		return
	}

	// email validator
	if emailErr := checkmail.ValidateFormat(data.Email); emailErr != nil {
		logger.Error(signupLogTag, requestID, "signup validate email failed, error : %+v", emailErr)
		DefaultResponse(c, http.StatusBadRequest, data, "Email is invalid (%s)", data.Email)
		c.Abort()
		return
	}

	// search if email already registered
	resEmail, errGetUserEmail := userModel.GetUserByEmail(data.Email)
	if resEmail.Email != "" {
		logger.Error(signupLogTag, requestID, "signup email already registered, email : %s", resEmail.Email)
		DefaultResponse(c, http.StatusForbidden, data, "Email %s already in use!", data.Email)
		c.Abort()
		return
	}

	if errGetUserEmail != nil {
		logger.Error(signupLogTag, requestID, "signup get user by email error, error : %+v", errGetUserEmail)
		DefaultResponse(c, http.StatusForbidden, data, "Email %s already in use!", data.Email)
		c.Abort()
		return
	}

	// Check if there was an error when saving user
	if errSignup := userModel.Signup(data); errSignup != nil {
		logger.Error(signupLogTag, requestID, "signup user, error : %+v", errSignup)
		DefaultResponse(c, http.StatusBadRequest, data, "Problem creating an account")
		c.Abort()
		return
	}

	logger.Info(signupLogTag, requestID, "New account registered (%s)", data.Email)
	DefaultResponse(c, http.StatusCreated, data, "New account registered (%s)", data.Email)
}

// Login allows a user to login a user and get
// access token
func (u *UserController) Login(c *gin.Context) {
	var data modelstruct.LoginUserCommand
	var tokenResult = &modelstruct.TokenResult{}
	requestID := requestid.Get(c)

	// Bind the request body data to var data and check if all details are provided
	if bindErr := c.BindJSON(&data); bindErr != nil {
		logger.Error(loginLogTag, requestID, "signup JSON bind failed, error : %+v", bindErr)
		DefaultResponse(c, http.StatusNotAcceptable, data, "Request provided is not acceptable")
		c.Abort()
		return
	}

	user, errEmail := userModel.GetUserByEmail(data.Email)

	if user.Email == "" {
		logger.Warn(loginLogTag, requestID, "User email %s account not found", data.Email)
		DefaultResponse(c, http.StatusNotFound, data, "User email %s account not found", data.Email)
		c.Abort()
		return
	}

	if errEmail != nil {
		logger.Error(loginLogTag, requestID, "Error on login sytem, error: %+v", errEmail.Error())
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
		logger.Error(loginLogTag, requestID, "Error on login sytem, error: %+v", errHash.Error())
		DefaultResponse(c, http.StatusForbidden, data, "Invalid user credentials")
		c.Abort()
		return
	}

	ts, err := helpers.CreateToken(user.ID.Hex())
	if err != nil {
		logger.Error(loginLogTag, requestID, "Error on login sytem, error: %+v", err.Error())
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := CreateAuth(user.ID.Hex(), ts)
	if saveErr != nil {
		logger.Error(loginLogTag, requestID, "Error on login sytem, error: %+v", saveErr.Error())
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}

	tokenResult.AccessToken = ts.AccessToken
	tokenResult.RefreshToken = ts.RefreshToken

	logger.Info(loginLogTag, requestID, "login success for email: %s", data.Email)
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
	requestID := requestid.Get(c)
	au, err := helpers.ExtractTokenMetadata(c.Request)
	if err != nil {
		logger.Error(logoutLogTag, requestID, "Error on logout system, error: %+v", err)
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	deleted, delErr := helpers.DeleteAuth(au.AccessUUID)
	if delErr != nil || deleted == 0 {
		logger.Error(logoutLogTag, requestID, "Error on logout system, error: %+v, deleted: %d", delErr, deleted)
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}
