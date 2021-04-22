package cmd

import (
	"fmt"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/spf13/cobra"
	"gitlab.xtc.home/xtc/redisearchd/conn/redis"
	"gitlab.xtc.home/xtc/redisearchd/http"
	"log"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start RediSearch Restful API",
	Long:  `Start RediSearch Restful API`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PreRun: func(cmd *cobra.Command, args []string) {
		raw := conf.String("redis.addr")
		redis.Init(raw)
	},
	Run: func(cmd *cobra.Command, args []string) {
		port := conf.Int("web.port")
		addr := fmt.Sprintf(":%d", port)
		http.Start(addr)
	},
	PostRun: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	flags := startCmd.PersistentFlags()

	flags.IntP("web.port", "", 16379, "web listening port")
	flags.StringP("redis.addr", "", "127.0.0.1:6379", "redis server addr")

	provider := posflag.Provider(flags, ".", conf)
	if err := conf.Load(provider, nil); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
}
