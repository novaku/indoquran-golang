package modelstruct

import "gopkg.in/mgo.v2/bson"

// User defines user object structure
type User struct {
	ID         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	Email      string        `json:"email" bson:"email"`
	Password   string        `json:"password" bson:"password"`
	IsVerified bool          `json:"is_verified" bson:"is_verified"`
}

// SignupUserCommand defines user form struct
type SignupUserCommand struct {
	// binding:"required" ensures that the field is provided
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginUserCommand defines user login form struct
type LoginUserCommand struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Token : token generated bucket
type Token struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

// TokenDetails : token user detail
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

// TokenResult : tokern output for controller
type TokenResult struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// AccessDetails : user access detail
type AccessDetails struct {
	AccessUUID string
	UserID     uint64
}
