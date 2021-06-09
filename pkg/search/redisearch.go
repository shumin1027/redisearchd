package search

import (
	"github.com/RediSearch/redisearch-go/redisearch"
	"gitlab.xtc.home/xtc/redisearchd/conn/redis"
)

var (
	connPool  redisearch.ConnPool
	searchMap map[string]*redisearch.Client
)

func init() {
	searchMap = make(map[string]*redisearch.Client)
}

func InitPool(address, password string) {
	connPool = redisearch.NewSingleHostPool(address)
}

func NewClient(name string) (c *redisearch.Client) {
	v, ok := searchMap[name]
	if ok {
		return v
	}
	c = redisearch.NewClientFromPool(redis.Pool(), name)
	searchMap[name] = c
	return c
}

func NewConnPool() redisearch.ConnPool {
	return connPool
}
