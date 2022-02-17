package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.xtc.home/xtc/redisearchd/app/utility"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall",
	Long:  `Uninstall`,
	Run: func(cmd *cobra.Command, args []string) {
		utility.UnInstall(conf.String("install.top"))
	},
}

func init() {
	flags := rootCmd.PersistentFlags()

	flags.StringP("install.top", "", "/usr/local/clustermom", "clustermom install path ")
	rootCmd.AddCommand(uninstallCmd)

}
