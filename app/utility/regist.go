package utility

import (
	"fmt"
	"time"

	consul "github.com/hashicorp/consul/api"
	"github.com/shirou/gopsutil/v3/host"
	"gitlab.xtc.home/xtc/redisearchd/app"
)

const (
	defaultCheckIntervalDuration = time.Duration(10) * time.Second
)

func Regist(agent *consul.Agent, ip string, port int, tags []string, meta map[string]string) {
	hostInfo, _ := host.Info()

	if tags == nil {
		tags = []string{}
	}

	tags = append(tags, hostInfo.Hostname)

	tags = append(tags, hostInfo.HostID)

	if meta == nil {
		meta = make(map[string]string)
	}

	meta["version"] = app.Version

	interval := defaultCheckIntervalDuration
	// deregister := time.Duration(1) * time.Minute

	service := &consul.AgentServiceRegistration{
		ID:      app.Name + "@" + hostInfo.HostID, // 服务节点UUID
		Name:    app.Name,                         // 服务名称 此处根据prometheus.yml配置进行约定
		Address: ip,                               // 服务 IP
		Port:    port,                             // 服务端口
		Tags:    tags,                             // tag，可以为空
		Meta:    meta,                             // meta，可以为空
		Check: &consul.AgentServiceCheck{ // 健康检查
			Interval: interval.String(), // 健康检查间隔
			//TCP:                            fmt.Sprintf("%s:%d", local.IP, port),             // 使用TCP做健康检查
			HTTP: fmt.Sprintf("http://%s:%d/ping", ip, port), // 使用HTTP做健康检查
			//DeregisterCriticalServiceAfter: deregister.String(),                              // 注销时间，相当于过期时间
		},
	}

	if err := agent.ServiceRegister(service); err != nil {
		panic(err)
	}
}
