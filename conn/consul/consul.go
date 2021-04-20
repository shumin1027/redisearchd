package consul

import (
	"errors"
	consulapi "github.com/hashicorp/consul/api"
	"gitlab.xtc.home/xtc/redisearchd/pkg/log"
	"go.uber.org/zap"
	"net/url"
)

type Consul struct {
	url    string
	client *consulapi.Client
	config *consulapi.Config
}

var (
	consul = new(Consul)
)

func Init(raw string) *consulapi.Client {
	consul.url = raw
	return Connect(parseConsulURL(raw))
}

func parseConsulURL(raw string) (scheme, address, dc string) {
	u, _ := url.Parse(raw)
	dc = u.Query().Get("dc")
	scheme = u.Scheme
	address = u.Host
	return
}

func Connect(scheme, address, dc string) *consulapi.Client {
	consul.config = &consulapi.Config{
		Datacenter: dc,
		Address:    address,
		Scheme:     scheme,
	}
	return connect(consul.config)
}

func connect(cfg *consulapi.Config) *consulapi.Client {
	consul.client, _ = consulapi.NewClient(cfg)
	return consul.client
}

func Client() *consulapi.Client {
	if consul == nil || consul.client == nil {
		if consul.config != nil {
			consul.client = connect(consul.config)
		} else {
			panic("place init consul first")
		}
	}
	return consul.client
}

func KVGet(key string) ([]byte, error) {
	client := Client()
	kv := client.KV()
	kvp, _, err := kv.Get(key, nil)
	if err != nil {
		log.Error("read consul kv error", zap.String("key", key), zap.Error(err))
		return nil, err
	}
	if kvp == nil {
		log.Debug("the key does not exist", zap.String("key", key))
		return nil, errors.New("the key does not exist")
	}
	return kvp.Value, err
}
