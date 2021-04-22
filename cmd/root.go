package cmd

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/spf13/cobra"
	"gitlab.xtc.home/xtc/redisearchd/app"
	"log"
	"os"
)

var conf = koanf.New(".")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: app.BuildInfo(),
	Use:     app.Name,
	Short:   "RediSearch Restful API",
	Long:    `RediSearch Restful API`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		provider := posflag.Provider(cmd.PersistentFlags(), ".", conf)
		if err := conf.Load(provider, nil); err != nil {
			log.Fatalf("error loading config: %v", err)
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		startCmd.PreRun(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		startCmd.Run(cmd, args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Panic(err)
		os.Exit(1)
	}
}
