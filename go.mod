module gitlab.xtc.home/xtc/redisearchd

go 1.15

require (
	github.com/RediSearch/redisearch-go v1.0.2-0.20201130114103-3264ad8d2487
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/arsmn/fiber-swagger/v2 v2.6.0
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/gofiber/fiber/v2 v2.7.1
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/iancoleman/strcase v0.1.2
	github.com/json-iterator/go v1.1.10
	github.com/klauspost/compress v1.11.13 // indirect
	github.com/magiconair/properties v1.8.4 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/mapstructure v1.4.0 // indirect
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/spf13/afero v1.5.1 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.1.1
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.7.1
	github.com/swaggo/swag v1.7.0
	github.com/valyala/fasthttp v1.23.0 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20210403161142-5e06dd20ab57 // indirect
	golang.org/x/text v0.3.6 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
)

replace github.com/RediSearch/redisearch-go => ./libs/redisearch-go
