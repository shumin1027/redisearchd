package search

import (
	"github.com/RediSearch/redisearch-go/redisearch"
	"gitlab.xtc.home/xtc/redisearchd/pkg/log"
	"go.uber.org/zap"
	"net/url"
	"strconv"
)

var (
	connPool  *redisearch.SingleHostPool
	searchMap map[string]*redisearch.Client
)

func init() {
	searchMap = make(map[string]*redisearch.Client)
}

func InitPool(raw string) {
	address, _, _ := ParseRedisURL(raw)
	connPool = redisearch.NewSingleHostPool(address)
}

func NewClient(name string) (c *redisearch.Client) {
	v, ok := searchMap[name]
	if ok {
		return v
	}
	c = redisearch.NewClientFromPool(connPool.Pool, name)
	searchMap[name] = c
	return c
}

func NewConnPool() redisearch.ConnPool {
	return connPool
}

//func NewRedisConn() redis.Conn {
//	return connPool.Get()
//}

func Close() error {
	return connPool.Close()
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
