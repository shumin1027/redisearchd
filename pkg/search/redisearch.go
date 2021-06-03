package search

import (
	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/gomodule/redigo/redis"
)

var (
	pool      *redis.Pool
	connPool  redisearch.ConnPool
	searchMap map[string]*redisearch.Client
)

func init() {
	searchMap = make(map[string]*redisearch.Client)
}

func Close() error {
	return pool.Close()
}

func NewPool(address, password string) {
	connPool = redisearch.NewSingleHostPool(address)

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

func NewClient(name string) (c *redisearch.Client) {
	v, ok := searchMap[name]
	if ok {
		return v
	}
	c = redisearch.NewClientFromPool(pool, name)
	searchMap[name] = c
	return c
}

func NewConnPool() redisearch.ConnPool {
	return connPool
}
