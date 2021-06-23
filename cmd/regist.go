package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"gitlab.xtc.home/xtc/redisearchd/app/utility"
	"gitlab.xtc.home/xtc/redisearchd/conn/consul"
	"gitlab.xtc.home/xtc/redisearchd/pkg/utils"
)

// registCmd represents the regist command
var registCmd = &cobra.Command{
	Use:   "regist",
	Short: "Regist Service To Consul",
	Long:  `Regist Service To Consul`,
	Run: func(cmd *cobra.Command, args []string) {
		port := conf.Int("http.port")

		ip := conf.String("ip")
		if ip == "" {
			ip = utils.ResolveDefaultIP()
		}

		url := conf.String("consul.url")
		client := consul.Init(url)

		tags := conf.Strings("consul.tags")

		meta := make(map[string]string)
		metas := conf.Strings("consul.meta")
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
	flags.IntP("http.port", "p", 16379, "server port")
	flags.StringP("consul.url", "u", "http://consul.service.consul:8500", "consul server url")
	flags.StringSliceP("consul.tags", "t", []string{}, "consul service tags, e.g: --consul.tags=t1,t2 --consul.tags=t3")
	flags.StringSliceP("consul.meta", "m", []string{}, "consul service meta, e.g: --consul.meta=key:value")
}
