package redis

import (
	"context"
	"github.com/gomodule/redigo/redis"
	conn "gitlab.xtc.home/xtc/redisearchd/conn/redis"
	"gitlab.xtc.home/xtc/redisearchd/pkg/log"
)

type PipeLine struct {
	redis.Conn
}

func Pipeline(ctx context.Context) *PipeLine {
	return &PipeLine{conn.Client()}
}

func (p *PipeLine) HSet(ctx context.Context, key string, values map[string]interface{}) error {
	args := redis.Args{}

	args = args.Add(key)
	for k, v := range values {
		args = args.Add(k)
		args = args.Add(v)
	}

	err := p.Send("HSET", args...)
	_ = p.Flush()

	return err
}

func (p *PipeLine) Del(ctx context.Context, keys ...string) error {
	args := redis.Args{}
	for _, key := range keys {
		args = args.Add(key)
	}
	err := p.Send("DEL", args...)

	return err
}

func (p *PipeLine) Exec(ctx context.Context) (interface{}, error) {
	_ = p.Flush()
	do, err := p.Receive()
	if err != nil {
		log.Warn("redis pipeline exec error")
		return do, err
	}
	return do, err
}
