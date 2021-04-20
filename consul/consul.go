package consul

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"net/url"
)

var config *consulapi.Config
var consul *consulapi.Client

func parseConsulURL(raw string) (scheme, address, dc string) {
	u, _ := url.Parse(raw)
	dc = u.Query().Get("dc")
	scheme = u.Scheme
	address = u.Host
	return
}

func Init(raw string) *consulapi.Client {
	return Connect(parseConsulURL(raw))
}

func Connect(scheme, address, dc string) *consulapi.Client {
	config = &consulapi.Config{
		Datacenter: dc,
		Address:    address,
		Scheme:     scheme,
	}
	return connect(config)
}

func connect(cfg *consulapi.Config) *consulapi.Client {
	consul, _ = consulapi.NewClient(cfg)
	return consul
}

func Consul() *consulapi.Client {
	if consul == nil {
		if config != nil {
			consul = connect(config)
		} else {
			panic("place init consul first")
		}
	}
	return consul
}

/*
根据ServiceName 和 Tags 查找服务
*/
func LookupService(name string, tags ...string) map[string]*consulapi.AgentService {
	taglen := len(tags)
	consul := Consul()
	agent := consul.Agent()

	filter := fmt.Sprintf("Service == \"%s\" ", name)
	if taglen > 0 {
		for _, tag := range tags {
			if len(tag) > 0 {
				filter = fmt.Sprintf("%s and Tags contains \"%s\" ", filter, tag)
			}
		}
	}
	services, err := agent.ServicesWithFilter(filter)
	if err != nil {
		println(err.Error())
	}
	return services
}
