package services

import (
	"indoquran-golang/config"
	"indoquran-golang/config/static"

	"github.com/go-redis/redis/v7"
)

// SetRedis : setup for redis connection
func SetRedis() *redis.Client {
	//Initializing redis
	dsn := config.LoadConfig().Redis.Host
	if len(dsn) == 0 {
		dsn = static.RedisDefaultHost
	}
	client := redis.NewClient(&redis.Options{
		Addr: dsn, //redis host and port
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return client
}
