package redis

import (
	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	pool    *redis.Pool
	clients map[string]*redisearch.Client
)

func init() {
	clients = make(map[string]*redisearch.Client)
}

func Init(url string) *redis.Pool {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(url)
		},
	}
	pool.TestOnBorrow = func(c redis.Conn, t time.Time) (err error) {
		if time.Since(t) > time.Second {
			_, err = c.Do("PING")
		}
		return err
	}
	return pool
}

func Client(name string) (c *redisearch.Client) {
	v, ok := clients[name]
	if ok {
		return v
	}
	c = redisearch.NewClientFromPool(pool, name)
	clients[name] = c
	return c
}

func DummyClient() *redisearch.Client {
	return Client("_redisearchd_")
}

func Pool() *redis.Pool {
	return pool
}

func Close() error {
	return pool.Close()
}
