package redis

import (
	"strings"

	"github.com/RediSearch/redisearch-go/redisearch"
)

var addr string
var pool redisearch.ConnPool
var clients map[string]*redisearch.Client

func init() {
	clients = make(map[string]*redisearch.Client)
}

func Init(server string) redisearch.ConnPool {
	addr = server
	return ConnPool()
}

func ConnPool() redisearch.ConnPool {
	addrs := strings.Split(addr, ",")
	if pool == nil {
		if len(addrs) == 1 {
			pool = redisearch.NewSingleHostPool(addrs[0])
		} else {
			pool = redisearch.NewMultiHostPool(addrs)
		}
	}
	return pool
}

func Client(index string) *redisearch.Client {
	if client, ok := clients[index]; ok {
		return client
	}
	client := redisearch.NewClient(addr, index)
	clients[index] = client
	return client
}

func DummyClient() *redisearch.Client {
	return Client("_redisearchd_")
}

func Close() error {
	return pool.Close()
}
