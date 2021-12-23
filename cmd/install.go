package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.xtc.home/xtc/redisearchd/app/utility"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install as a systemd unit",
	Long:  `Install as a systemd unit`,
	Run: func(cmd *cobra.Command, args []string) {
		top := conf.String("install.top")
		utility.Install(top)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	flags := installCmd.PersistentFlags()
	flags.StringP("install.top", "", "/usr/local/clustermom", "install path")
}
