package tests

import (
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
	"testing"
)

func TestRedisearchConn(t *testing.T) {
	connPool := redisearch.NewSingleHostPool("redis://16.16.18.249:6379/0")
	do, err := connPool.Get().Do("ping")
	fmt.Println(do)
	fmt.Println(err)
}
