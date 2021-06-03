package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.xtc.home/xtc/redisearchd/conn/redis"
	"gitlab.xtc.home/xtc/redisearchd/http"
	"gitlab.xtc.home/xtc/redisearchd/pkg/log"
	"gitlab.xtc.home/xtc/redisearchd/pkg/search"
	"go.uber.org/zap"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start RediSearch Restful API",
	Long:  `Start RediSearch Restful API`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PreRun: func(cmd *cobra.Command, args []string) {
		raw := conf.String("redis.url")

		address, password, _ := redis.ParseRedisURL(raw)
		search.NewPool(address, password)
		err := redis.InitRedis(raw)
		if err != nil {
			log.Logger().Fatal("redis client init error", zap.Error(err))
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		port := conf.Int("http.port")
		addr := fmt.Sprintf(":%d", port)
		http.Start(addr)
	},
	PostRun: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	flags := startCmd.PersistentFlags()

	flags.IntP("http.port", "", 16379, "web listening port")
	flags.StringP("redis.url", "", "redis://redis.service.consul:6379?db=0", "redis server url")

}
