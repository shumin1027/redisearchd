package redisearch

import (
	"context"
	"encoding/json"
	"github.com/RediSearch/redisearch-go/redisearch"
	"gitlab.xtc.home/xtc/redisearchd/conn/search"
	"testing"
)

func TestCreateIndex(t *testing.T) {
	cli := search.NewClient("test")
	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("body")).
		AddField(redisearch.NewTextFieldOptions("title", redisearch.TextFieldOptions{Weight: 5.0, Sortable: true})).
		AddField(redisearch.NewNumericField("date"))

	def := redisearch.NewIndexDefinition().AddPrefix("redisearch:")
	CreateIndex(context.Background(), cli, sc, def)
}

func TestListIndexes(t *testing.T) {
	cli := search.NewClient("test")
	indexes, _ := ListIndexes(context.Background(), cli)
	data, _ := json.Marshal(indexes)
	println(string(data))
}

func TestInfo(t *testing.T) {
	cli := search.NewClient("test")
	info, _ := Info(context.Background(), cli)
	data, _ := json.Marshal(info)
	println(string(data))
}
