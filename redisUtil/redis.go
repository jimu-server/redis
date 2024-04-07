package redisUtil

import (
	"config/config"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Evn.App.Redis.Host, config.Evn.App.Redis.Port),
		Password: config.Evn.App.Redis.Password, // no password set
		DB:       config.Evn.App.Redis.DB,       // use default DB
	})
}

func Set(key string, value interface{}) error {
	return nil
}
