package cache

import (
	"LibraryApi/internal/config"
	"errors"
	"github.com/go-redis/redis/v7"
	"strconv"
)

var redisClient *redis.Client

func InitializeRedis(config config.CacheConfig) error {
	db, err := strconv.Atoi(config.DbName)
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       db,
	})
	_, err = redisClient.Ping().Result()
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func GetRedisClient() *redis.Client {
	return redisClient
}
