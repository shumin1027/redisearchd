package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gitlab.xtc.home/xtc/redisearchd/pkg/log"
	"go.uber.org/zap"
	"net/url"
	"strconv"
)

var redisClient redis.UniversalClient

func InitRedis(url string) error {
	address, password, database := ParseRedisURL(url)

	redisClient = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        []string{address},
		Password:     password,
		DB:           database,
		MaxRetries:   5,
		MinIdleConns: 10,
		PoolSize:     1000,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Logger().Error("Init Redis Error", zap.Error(err))
		return err
	}
	return nil
}

func Client() redis.UniversalClient {
	return redisClient
}

func Close() error {
	return redisClient.Close()
}

func ParseRedisURL(raw string) (address, password string, database int) {
	u, err := url.Parse(raw)
	if err != nil {
		log.Logger().Warn("parse redis raw failed", zap.Error(err))
	}
	address = u.Host

	db := u.Path[1:len(u.Path)]
	database, err = strconv.Atoi(db)
	if err != nil {
		log.Logger().Warn("parse redis database failed", zap.Error(err))
	}

	password = u.Query().Get("password")
	return address, password, database
}
