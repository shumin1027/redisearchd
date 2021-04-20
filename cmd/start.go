package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.xtc.home/xtc/redisearchd/conn/redis"
	"gitlab.xtc.home/xtc/redisearchd/http"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start RediSearch Restful API",
	Long:  `Start RediSearch Restful API`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PreRun: func(cmd *cobra.Command, args []string) {
		raw := viper.GetString("redis.addr")
		redis.Init(raw)
	},
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetInt("web.port")
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
	viper.BindPFlag("web.port", flags.Lookup("web.port"))

	flags.StringP("redis.addr", "", "127.0.0.1:6379", "redis server addr")
	viper.BindPFlag("redis.addr", flags.Lookup("redis.addr"))
}
