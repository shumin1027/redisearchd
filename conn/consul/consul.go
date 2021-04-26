package consul

import (
	consulapi "github.com/hashicorp/consul/api"
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
