package redis

import (
	"context"
	"github.com/gomodule/redigo/redis"
	conn "gitlab.xtc.home/xtc/redisearchd/conn/redis"
)

func Keys(ctx context.Context, pattern string) ([]string, error) {
	var keys []string
	values, err := redis.Values(conn.Client().Do("KEYS", pattern))
	if err != nil {
		return keys, err
	}
	for _, value := range values {
		v := string(value.([]byte))
		keys = append(keys, v)
	}
	return keys, err
}

func Del(ctx context.Context, keys ...string) (int64, error) {
	args := redis.Args{}
	for _, key := range keys {
		args = args.Add(key)
	}
	v, err := conn.Client().Do("DEL", args...)
	num := v.(int64)
	return num, err
}

func HSet(ctx context.Context, key string, values map[string]interface{}) (interface{}, error) {
	args := redis.Args{}

	args = args.Add(key)
	for k, v := range values {
		args = args.Add(k)
		args = args.Add(v)
	}

	v, err := conn.Client().Do("HSET", args...)
	return v, err
}
