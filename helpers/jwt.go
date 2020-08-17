package helpers

import (
	"indoquran-golang/config"
	"indoquran-golang/models/modelstruct"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var serverSecret = config.LoadConfig().Server.Secret
var jwtKey = []byte(serverSecret)

// Claims defines jwt claims
type Claims struct {
	UserID string `json:"email"`
	jwt.StandardClaims
}

// CreateToken : create token jwt
func CreateToken(userid string) (*modelstruct.TokenDetails, error) {
	authSecret := config.LoadConfig().Auth
	td := &modelstruct.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * time.Duration(authSecret.AccessExpire)).Unix()
	td.AccessUUID = uuid.New().String()

	td.RtExpires = time.Now().Add(time.Hour * time.Duration(authSecret.RefreshExpire)).Unix()
	td.RefreshUUID = uuid.New().String()

	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(authSecret.AccessSecret))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(authSecret.RefreshSecret))
	if err != nil {
		return nil, err
	}
	return td, nil
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
