module gitlab.xtc.home/xtc/redisearchd

go 1.16

require (
	github.com/RediSearch/redisearch-go v1.1.1-0.20210416071559-f79df23649c6
	github.com/StackExchange/wmi v0.0.0-20210224194228-fe8f1750fd46 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/armon/go-metrics v0.3.6 // indirect
	github.com/arsmn/fiber-swagger/v2 v2.6.0
	github.com/coreos/go-systemd/v22 v22.3.1
	github.com/fatih/color v1.10.0 // indirect
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/gofiber/fiber/v2 v2.8.0
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/hashicorp/consul/api v1.8.1
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v0.16.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/iancoleman/strcase v0.1.3
	github.com/json-iterator/go v1.1.10
	github.com/klauspost/compress v1.12.1 // indirect
	github.com/knadh/koanf v0.16.0
	github.com/mitchellh/copystructure v1.1.2 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/pelletier/go-toml v1.9.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shirou/gopsutil v3.21.3+incompatible
	github.com/spf13/cobra v1.1.3
	github.com/swaggo/swag v1.7.0
	github.com/tklauser/go-sysconf v0.3.5 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
	golang.org/x/net v0.0.0-20210421230115-4e50805a0758 // indirect
	golang.org/x/sys v0.0.0-20210421221651-33663a62ff08 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

//replace github.com/RediSearch/redisearch-go => ./libs/redisearch-go
