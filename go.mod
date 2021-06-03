module gitlab.xtc.home/xtc/redisearchd

go 1.16

require (
	github.com/RediSearch/redisearch-go v1.1.1
	github.com/StackExchange/wmi v0.0.0-20210224194228-fe8f1750fd46 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/armon/go-metrics v0.3.8 // indirect
	github.com/arsmn/fiber-swagger/v2 v2.6.0
	github.com/coreos/go-systemd/v22 v22.3.2
	github.com/fatih/color v1.11.0 // indirect
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-redis/redis/v8 v8.9.0
	github.com/gofiber/fiber/v2 v2.10.0
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/hashicorp/consul/api v1.8.1
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v0.16.1 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/iancoleman/strcase v0.1.3
	github.com/json-iterator/go v1.1.11
	github.com/knadh/koanf v1.0.0
	github.com/pelletier/go-toml v1.9.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shirou/gopsutil v3.21.4+incompatible
	github.com/spf13/cobra v1.1.3
	github.com/swaggo/swag v1.7.0
	github.com/tklauser/go-sysconf v0.3.6 // indirect
	github.com/valyala/fasthttp v1.25.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20210520170846-37e1c6afe023 // indirect
	golang.org/x/tools v0.1.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

//replace github.com/RediSearch/redisearch-go => ./libs/redisearch-go
