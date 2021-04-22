package utility

import (
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"github.com/shirou/gopsutil/host"
	"gitlab.xtc.home/xtc/redisearchd/app"
	"time"
)

func Regist(agent *consul.Agent, ip string, port int, tags []string, meta map[string]string) {

	host, _ := host.Info()
	if tags == nil {
		tags = []string{}
	}
	tags = append(tags, host.Hostname)
	tags = append(tags, host.HostID)

	if meta == nil {
		meta = make(map[string]string)
	}
	meta["version"] = app.Version

	interval := time.Duration(10) * time.Second
	//deregister := time.Duration(1) * time.Minute

	service := &consul.AgentServiceRegistration{
		ID:      app.Name + "@" + host.HostID, // 服务节点UUID
		Name:    app.Name,                     // 服务名称 此处根据prometheus.yml配置进行约定
		Address: ip,                           // 服务 IP
		Port:    port,                         // 服务端口
		Tags:    tags,                         // tag，可以为空
		Meta:    meta,                         // meta，可以为空
		Check: &consul.AgentServiceCheck{ // 健康检查
			Interval: interval.String(), // 健康检查间隔
			//TCP:                            fmt.Sprintf("%s:%d", local.IP, port),             // 使用TCP做健康检查
			HTTP: fmt.Sprintf("http://%s:%d/ping", ip, port), // 使用HTTP做健康检查
			//DeregisterCriticalServiceAfter: deregister.String(),                              // 注销时间，相当于过期时间
		},
	}
	agent.ServiceRegister(service)
}
