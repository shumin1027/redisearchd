package redis

import (
	"github.com/gomodule/redigo/redis"
	"gitlab.xtc.home/xtc/redisearchd/pkg/log"
	"go.uber.org/zap"
	"net/url"
	"strconv"
)

var pool *redis.Pool

//var redisClient redis.UniversalClient

func InitRedis(raw string) {
	address, password, _ := ParseRedisURL(raw)
	pool = &redis.Pool{Dial: func() (redis.Conn, error) {
		var conn redis.Conn
		var err error
		if password != "" {
			conn, err = redis.Dial("tcp", address, redis.DialPassword(password))
		} else {
			conn, err = redis.Dial("tcp", address)
		}
		return conn, err
	}}
}

func Client() redis.Conn {
	return pool.Get()
}

func Pool() *redis.Pool {
	return pool
}

func Close() error {
	return pool.Close()
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
