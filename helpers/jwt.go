package helpers

import (
	"indoquran-golang/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var serverSecret = config.LoadConfig().Server.Secret
var jwtKey = []byte(serverSecret)

// Claims defines jwt claims
type Claims struct {
	UserID string `json:"email"`
	jwt.StandardClaims
}

// GenerateToken handles generation of a jwt code
// @returns string -> token and error -> err
func GenerateToken(userID string) (string, error) {
	var err error

	// Define token expiration time
	sessionExpire := config.LoadConfig().Session.Expire
	expirationTime := time.Now().Add(time.Minute * time.Duration(sessionExpire)).Unix()
	// Define the payload and exp time
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key encoding
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

// DecodeToken handles decoding a jwt token
func DecodeToken(tkStr string) (string, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tkStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !tkn.Valid {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
		return "", err
	}

	return claims.UserID, nil
}
