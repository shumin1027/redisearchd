package redisearch

import (
	"context"
	"github.com/RediSearch/redisearch-go/redisearch"
)

func Search(ctx context.Context, client *redisearch.Client, query *redisearch.Query) (docs []redisearch.Document, total int, err error) {
	return client.Search(query)
}
