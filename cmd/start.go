package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.xtc.home/xtc/redisearchd/conn/redis"
	"gitlab.xtc.home/xtc/redisearchd/http"
	"gitlab.xtc.home/xtc/redisearchd/pkg/search"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start RediSearch Restful API",
	Long:  `Start RediSearch Restful API`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PreRun: func(cmd *cobra.Command, args []string) {
		raw := conf.String("redis.url")
		redis.InitRedis(raw)
		address, password, _ := redis.ParseRedisURL(raw)
		search.InitPool(address, password)
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

	flags.IntP("http.port", "p", 16379, "web listening port")
	flags.StringP("redis.url", "", "redis://redis.service.consul:6379/0", "redis server url")

}
