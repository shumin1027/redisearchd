package redisearch

import (
	"context"

	"github.com/RediSearch/redisearch-go/redisearch"
)

// Returns information and statistics on the index.
func Info(ctx context.Context, client *redisearch.Client) (*redisearch.IndexInfo, error) {
	return client.Info()
}

// Returns a list of all existing indexes.
func ListIndexes(ctx context.Context, client *redisearch.Client) ([]string, error) {
	return client.List()
}

func CreateIndex(ctx context.Context, client *redisearch.Client, schema *redisearch.Schema, definition *redisearch.IndexDefinition) error {
	return client.CreateIndexWithIndexDefinition(schema, definition)
}

func DropIndex(ctx context.Context, client *redisearch.Client, deldocs bool) error {
	return client.DropIndex(deldocs)
}
