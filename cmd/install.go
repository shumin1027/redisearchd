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
		utility.Install()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
