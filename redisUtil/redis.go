package redisUtil

import (
	"context"
	"fmt"
	"github.com/jimu-server/config"
	"github.com/jimu-server/logger"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"time"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Evn.App.Redis.Host, config.Evn.App.Redis.Port),
		Password: config.Evn.App.Redis.Password, // no password set
		DB:       config.Evn.App.Redis.DB,       // use default DB
	})
	ping := client.Ping(context.Background())
	if ping.Err() != nil {
		logger.Logger.Panic(ping.Err().Error())
	}

}

func Set(key string, value interface{}) error {
	return setSetEx(key, value, 0)
}

func SetEx(key string, value any, expiration time.Duration) error {
	return setSetEx(key, value, expiration)
}

func Get(key string) (string, error) {
	return get(key)
}

func setSetEx(key string, value any, expiration time.Duration) error {
	var buffer []byte
	v := ""
	switch value.(type) {
	case string:
		v = value.(string)
	case int:
	default:
		buffer, _ = jsoniter.Marshal(value)
		v = string(buffer)
	}
	cmd := client.SetEx(context.Background(), key, v, expiration)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func get(key string) (string, error) {
	cmd := client.Get(context.Background(), key)
	if cmd.Err() != nil {
		return "", cmd.Err()
	}
	return cmd.Val(), nil
}
