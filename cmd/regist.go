package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.xtc.home/xtc/redisearchd/app/utility"
	consul2 "gitlab.xtc.home/xtc/redisearchd/conn/consul"
	"gitlab.xtc.home/xtc/redisearchd/pkg/utils"
	"strings"
)

// registCmd represents the regist command
var registCmd = &cobra.Command{
	Use:   "regist",
	Short: "Regist Service To Consul",
	Long:  `Regist Service To Consul`,
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetInt("web.port")

		ip := viper.GetString("ip")
		if ip == "" {
			ip = utils.ResolveDefaultIP()
		}

		url := viper.GetString("consul.url")

		client := consul2.Init(url)

		tags := viper.GetStringSlice("consul.tags")
		println(tags)

		meta := make(map[string]string)
		metas := viper.GetStringSlice("consul.meta")
		for _, m := range metas {
			v := strings.Split(m, ":")
			key := v[0]
			value := v[1]
			meta[key] = value
		}

		utility.Regist(client.Agent(), ip, port, tags, meta)
	},
}

func init() {
	rootCmd.AddCommand(registCmd)
	flags := registCmd.PersistentFlags()

	flags.StringP("ip", "i", "", "server ip")
	viper.BindPFlag("ip", flags.Lookup("ip"))

	flags.IntP("port", "p", 13000, "server port")
	_ = viper.BindPFlag("port", flags.Lookup("port"))

	flags.StringP("consul.url", "u", "http://consul.service.consul:8500", "consul server url")
	viper.BindPFlag("consul.url", flags.Lookup("consul.url"))

	flags.StringSliceP("consul.tags", "t", []string{}, "consul service tags, e.g: --consul.tags=t1,t2 --consul.tags=t3")
	viper.BindPFlag("consul.tags", flags.Lookup("consul.tags"))

	flags.StringSliceP("consul.meta", "m", []string{}, "consul service meta, e.g: --consul.meta=key:value")
	viper.BindPFlag("consul.tags", flags.Lookup("consul.tags"))

}
