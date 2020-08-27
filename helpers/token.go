package helpers

import (
	"fmt"
	"indoquran-golang/config"
	"indoquran-golang/models/modelstruct"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v7"
)

var (
	client     *redis.Client
	currentDir string
)

const (
	extractTokenMetadataLogTag = "helpers|token.go|ExtractTokenMetadata()"
)

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

	currentDir = os.Args[0]
}

// TokenValid : check token is valid
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// VerifyToken : verify the token
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.LoadConfig().Auth.AccessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ExtractToken : extract the token info
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normaly Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// ExtractTokenMetadata : extract meta data from token
func ExtractTokenMetadata(r *http.Request, requestID string) (*modelstruct.AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}

		return &modelstruct.AccessDetails{
			AccessUUID: accessUUID,
			UserID:     fmt.Sprintf("%s", claims["user_id"]),
		}, nil
	}
	return nil, err
}

// DeleteAuth : delete auth token user
func DeleteAuth(givenUUID string) (int64, error) {
	deleted, err := client.Del(givenUUID).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
