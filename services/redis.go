package services

import (
	"bitbucket.org/indoquran-api/config"

	"github.com/go-redis/redis/v7"
)

// SetRedis : setup for redis connection
func SetRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: config.LoadConfig().Redis.Host + ":" + config.LoadConfig().Redis.Port, //redis host and port
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return client
}
