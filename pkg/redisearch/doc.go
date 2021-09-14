package redisearch

import (
	"context"
	"errors"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/gomodule/redigo/redis"
	"gitlab.xtc.home/xtc/redisearchd/pkg/log"
	"go.uber.org/zap"
)

type Document struct {
	redisearch.Document
}
type DocumentList []Document

func NewDocument(id string, score float32) Document {
	doc := redisearch.Document{
		Id:         id,
		Score:      score,
		Properties: make(map[string]interface{}),
	}
	return Document{
		doc,
	}
}

// SetPayload Sets the document payload
func (d *Document) SetPayload(payload []byte) {
	d.Payload = payload
}

// Set sets a property and its value in the document
func (d Document) Set(name string, value interface{}) Document {
	d.Properties[name] = value
	return d
}

// SetPayload Sets the document payload
func (d *Document) loadFields(lst []interface{}) *Document {
	for i := 0; i < len(lst); i += 2 {
		var prop string
		switch lst[i].(type) {
		case []byte:
			prop = string(lst[i].([]byte))
		default:
			prop = lst[i].(string)
		}

		var val interface{}
		switch v := lst[i+1].(type) {
		case []byte:
			val = string(v)
		default:
			val = v
		}
		*d = d.Set(prop, val)
	}
	return d
}

func (doc *Document) Serialize(args redis.Args) redis.Args {
	args = append(args, doc.Id)
	for k, f := range doc.Properties {
		args = append(args, k, f)
	}
	return args
}

func AddDocs(ctx context.Context, connpool redisearch.ConnPool, docs ...Document) error {
	conn := connpool.Get()
	defer conn.Close()

	n := 0
	var merr redisearch.MultiError

	for ii, doc := range docs {
		args := make(redis.Args, 0, 1+len(doc.Properties))
		args = doc.Serialize(args)

		if err := conn.Send("HSET", args...); err != nil {
			if merr == nil {
				merr = redisearch.NewMultiError(len(docs))
			}
			merr[ii] = err
			return merr
		}
		n++
	}

	if err := conn.Flush(); err != nil {
		return err
	}

	for n > 0 {
		reply, err := conn.Receive()
		println(reply)
		if err != nil {
			if merr == nil {
				merr = redisearch.NewMultiError(len(docs))
			}
			merr[n-1] = err
		}
		n--
	}

	if merr == nil {
		return nil
	}
	return merr
}

func GetDocById(ctx context.Context, connpool redisearch.ConnPool, id string, fields ...string) (*Document, error) {
	conn := connpool.Get()
	defer conn.Close()

	command := "HGETALL"
	args := redis.Args{id}
	if len(fields) > 1 {
		command = "HMGET"
		args.Add(fields)
	}
	reply, err := conn.Do(command, args...)
	if err != nil {
		return nil, err
	}
	var doc *Document
	if reply != nil {
		var array_reply []interface{}
		array_reply, err = redis.Values(reply, err)
		if err != nil {
			return nil, err
		}
		if len(array_reply) > 0 {
			document := NewDocument(id, 1)
			document.loadFields(array_reply)
			doc = &document
		}
	}
	return doc, nil
}

func DeleteDocs(ctx context.Context, connpool redisearch.ConnPool, ids ...string) error {
	conn := connpool.Get()
	defer conn.Close()

	args := redis.Args{}

	for _, id := range ids {
		args = args.Add(id)
	}

	if _, err := conn.Do("DEL", args...); err != nil {
		return err
	}

	return nil
}

func UpdateDocs(ctx context.Context, connpool redisearch.ConnPool, docs ...Document) error {
	conn := connpool.Get()
	defer conn.Close()

	for _, doc := range docs {
		args := make(redis.Args, 0, 1+len(doc.Properties))
		args = doc.Serialize(args)
		err := conn.Send("HSET", args...)
		if err != nil {
			return err
		}
	}

	if err := conn.Flush(); err != nil {
		return err
	}

	reply, err := conn.Receive()

	if err != nil {
		log.Logger().Error("update redis docs error", zap.Error(err), zap.Any("reply", reply))
	}
	return err
}

func UpdateDocFields(ctx context.Context, connpool redisearch.ConnPool, id string, keys []string, values []string) error {
	conn := connpool.Get()
	defer conn.Close()
	if len(keys) != len(values) {
		return errors.New("number of 'key' and 'value' is inconsistent")
	}
	args := redis.Args{id}
	for i, key := range keys {
		args = args.Add(key)
		args = args.Add(values[i])
	}
	err := conn.Send("HMSET", args...)
	if err != nil {
		return err
	}

	if err := conn.Flush(); err != nil {
		return err
	}

	reply, err := conn.Receive()

	if err != nil {
		log.Logger().Error("update redis docs error", zap.Error(err), zap.Any("reply", reply))
	}
	return err
}
