/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.xtc.home/xtc/redisearchd/conn"
	"gitlab.xtc.home/xtc/redisearchd/http"
	"os"
)

var cfgFile string

var startCmd = &cobra.Command{
	Use:   "redisearchd",
	Short: "RediSearch API",
	Long:  `RediSearch API`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PreRun: func(cmd *cobra.Command, args []string) {
		raw := viper.GetString("redis.addr")
		conn.Init(raw)
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
	cobra.OnInitialize(initConfig)

	self := startCmd

	self.PersistentFlags().IntP("web.port", "", 8080, "web listening port")
	viper.BindPFlag("web.port", rootCmd.PersistentFlags().Lookup("web.port"))

	self.PersistentFlags().StringP("redis.addr", "", "127.0.0.1:6379", "redis server addr")
	viper.BindPFlag("redis.addr", rootCmd.PersistentFlags().Lookup("redis.addr"))

	self.PersistentFlags().StringVarP(&cfgFile, "config", "c", "redisearchd.yaml", "config file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".redisearchd" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".redisearchd")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
