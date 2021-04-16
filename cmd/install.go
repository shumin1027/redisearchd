package cmd

import (
	"bufio"
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/coreos/go-systemd/v22/unit"
	"github.com/spf13/cobra"
	"gitlab.xtc.home/xtc/redisearchd/app"
	"gitlab.xtc.home/xtc/redisearchd/pkg/utils"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install as a systemd unit",
	Long:  `Install as a systemd unit`,
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.PersistentFlags().Lookup("bin-path").Value.String()
		unitFilePath := cmd.PersistentFlags().Lookup("unitfile-path").Value.String()

		binfile := filepath.Join(path, app.Name)
		log.Println("executable bin file:", binfile)

		unitfile := filepath.Join(unitFilePath, app.Name+".service")
		log.Println("systemd unit file:", unitfile)

		install(binfile, unitfile)
	},
}

func init() {
	self := installCmd
	rootCmd.AddCommand(self)
	self.PersistentFlags().StringP("bin-path", "b", "/usr/local/bin/", "default bin path")
	self.PersistentFlags().StringP("unitfile-path", "u", "/usr/lib/systemd/system/", "default systemd unit file path")
}

const unittext = `
[Unit]
Description=Redisearch Restful API
Documentation=http://gitlab.xtc.home/xtc/redisearchd.git
After=network.service

[Service]
Type=simple
Restart=on-failure
User=root
Group=root

[Install]
WantedBy=multi-user.target
`

func install(binfile, unitfile string) {

	if exists, _ := utils.PathExists(unitfile); exists {
		err := os.Remove(unitfile)
		if err != nil {
			log.Panic(err)
		}
	}

	if exists, _ := utils.PathExists(binfile); exists {
		err := os.Remove(binfile)
		if err != nil {
			log.Panic(err)
		}
	}

	self, err := os.Executable()
	if err != nil {
		log.Panic(err)
	}
	_, err = utils.CopyFile(self, binfile, 0751)
	if err != nil {
		log.Panic(err)
	}

	reader := strings.NewReader(unittext)
	opts, err := unit.Deserialize(reader)
	if err != nil {
		log.Panic(err)
	}
	opt := unit.NewUnitOption("Service", "ExecStart", binfile)
	opts = append(opts, opt)
	r := unit.Serialize(opts)

	var f *os.File
	exist, _ := utils.PathExists(unitfile)
	if exist {
		f, err = os.OpenFile(unitfile, os.O_APPEND, 0666)
		if err != nil {
			log.Panic(err)
		}
	} else {
		f, _ = utils.MakeFile(unitfile)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Panic(err)
		}
	}(f)
	w := bufio.NewWriter(f)
	_, err = w.ReadFrom(r)
	if err != nil {
		log.Panic(err)
	}
	defer func(w *bufio.Writer) {
		err := w.Flush()
		if err != nil {
			log.Panic(err)
		}
	}(w)

	ctx := context.TODO()
	systemd, err := dbus.NewWithContext(ctx)
	if err != nil {
		log.Panic(err)
	}
	defer systemd.Close()
	err = systemd.ReloadContext(ctx)
	if err != nil {
		log.Panic(err)
	}
}
