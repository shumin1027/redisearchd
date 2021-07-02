package tests

import (
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
	"gitlab.xtc.home/xtc/redisearchd/conn/redis"
	"strings"
	"testing"
)

const (
	field_tokenization = ",.<>{}[]\"':;!@#$%^&*()-+=~ "
)

func EscapeTextFileString(value string) string {
	for _, char := range field_tokenization {
		value = strings.Replace(value, string(char), ("\\" + string(char)), -1)
	}
	return value
}

func TestRedisearchConn(t *testing.T) {
	connPool := redisearch.NewSingleHostPool("127.0.0.1:6379")
	do, err := connPool.Get().Do("ping")
	fmt.Println(do)
	fmt.Println(err)

	redis.Init("redis://127.0.0.1:6379/0")
	cli := redis.Client("test")
	str := EscapeTextFileString("this is title")
	fmt.Println(str)

	q := &redisearch.Query{
		Raw: "@title:'this is title'",
		Paging: redisearch.Paging{
			Offset: 0,
			Num:    10,
		},
	}

	docs, total, err := cli.Search(q)
	if err != nil {
		println(err.Error())
	}
	println(total)
	fmt.Println(docs)
}
