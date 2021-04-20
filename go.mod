module gitlab.xtc.home/xtc/redisearchd

go 1.16

require (
	github.com/RediSearch/redisearch-go v1.1.0
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/arsmn/fiber-swagger/v2 v2.6.0
	github.com/coreos/go-systemd/v22 v22.3.1
	github.com/gofiber/fiber/v2 v2.7.1
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/hashicorp/consul/api v1.8.1
	github.com/iancoleman/strcase v0.1.3
	github.com/json-iterator/go v1.1.10
	github.com/klauspost/compress v1.12.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/shirou/gopsutil v3.21.3+incompatible
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/swaggo/swag v1.7.0
	github.com/valyala/fasthttp v1.23.0 // indirect
	gitlab.xtc.home/xtc/clustermom-agent v0.0.0-20210419073557-75ecd4bc6f95
	go.uber.org/zap v1.16.0
	golang.org/x/sys v0.0.0-20210419170143-37df388d1f33 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/RediSearch/redisearch-go => ./libs/redisearch-go
