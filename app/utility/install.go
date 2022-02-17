package utility

import (
	"bufio"
	"context"
	"fmt"
	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/coreos/go-systemd/v22/unit"
	"github.com/valyala/fasttemplate"
	"gitlab.xtc.home/xtc/redisearchd/app"
	"gitlab.xtc.home/xtc/redisearchd/pkg/log"
	"gitlab.xtc.home/xtc/redisearchd/pkg/utils"
	"os"
	"path/filepath"
	"strings"
)

var (
	Top                = "/usr/local/clustermom"
	BinPath            = "{{TOP}}/bin/"
	VarPath            = "{{TOP}}/var/"
	VarLibPath         = "{{TOP}}/var/lib/"
	EtcPath            = "{{TOP}}/etc/"
	UnitFilePath       = "{{TOP}}/etc/systemd/system/"
	SystemUnitFilePath = "/etc/systemd/system/"
)

const unittext = `
[Unit]
After=network.service

[Service]
Type=simple
Restart=on-failure
User=root
Group=root

[Install]
WantedBy=clustermom.target
`

func preInstall() {
	BinPath = filepath.Clean(fasttemplate.New(BinPath, "{{", "}}").ExecuteString(map[string]interface{}{
		"TOP": Top,
	}))
	VarPath = filepath.Clean(fasttemplate.New(VarPath, "{{", "}}").ExecuteString(map[string]interface{}{
		"TOP": Top,
	}))
	EtcPath = filepath.Clean(fasttemplate.New(EtcPath, "{{", "}}").ExecuteString(map[string]interface{}{
		"TOP": Top,
	}))
	VarLibPath = filepath.Clean(fasttemplate.New(VarLibPath, "{{", "}}").ExecuteString(map[string]interface{}{
		"TOP": Top,
	}))
	UnitFilePath = filepath.Clean(fasttemplate.New(UnitFilePath, "{{", "}}").ExecuteString(map[string]interface{}{
		"TOP": Top,
	}))

	utils.MakeDir(BinPath)
	utils.MakeDir(VarPath)
	utils.MakeDir(EtcPath)
	utils.MakeDir(VarLibPath)
	utils.MakeDir(UnitFilePath)
}

func Install(top string) {
	Top = top
	preInstall()
	log.Logger().Info("installing bin")
	installBin()
	log.Logger().Info("installing systemd unit")
	installUnit()
	log.Logger().Info("reloading systemd-daemon")
	ReloadSystemdDaemon()
	log.Logger().Info("install complete")
}

func UnInstall(top string) {
	Top = top
	binfile := filepath.Join(fasttemplate.New(BinPath, "{{", "}}").ExecuteString(map[string]interface{}{"TOP": Top}), app.Name)
	unitfile := filepath.Join(fasttemplate.New(UnitFilePath, "{{", "}}").ExecuteString(map[string]interface{}{"TOP": Top}), app.Name+".service")
	unitfileLink := filepath.Join(fasttemplate.New(SystemUnitFilePath, "{{", "}}").ExecuteString(map[string]interface{}{"TOP": Top}), fmt.Sprintf("clustermom-%s.service", app.Name))

	log.Info("uninstalling bin")
	os.Remove(binfile)
	log.Info("uninstalling systemd unit")
	os.Remove(unitfile)
	os.Remove(unitfileLink)
	log.Info("reloading systemd-daemon")
	ReloadSystemdDaemon()
	log.Info("uninstall complete")
}

func installBin() {
	binfile := filepath.Join(BinPath, app.Name)
	// rm old binfile
	if exists, _ := utils.Exists(binfile); exists {
		err := os.Remove(binfile)
		if err != nil {
			log.StdLogger().Panic(err)
		}
	}

	// install new binfile
	self, err := os.Executable()
	if err != nil {
		log.StdLogger().Panic(err)
	}
	_, err = utils.CopyFile(self, binfile, 0751)
	if err != nil {
		log.StdLogger().Panic(err)
	}
}

func installUnit() {
	unitfile := filepath.Join(UnitFilePath, app.Name+".service")
	// 链接到 /usr/lib/systemd/system/ 的时候加上前缀clustermom- 方便后期查找
	unitfile_link := filepath.Join(SystemUnitFilePath, fmt.Sprintf("clustermom-%s.service", app.Name))

	// rm old unitfile
	if exists, _ := utils.Exists(unitfile); exists {
		err := os.Remove(unitfile)
		if err != nil {
			log.StdLogger().Panic(err)
		}
	}
	if exists, _ := utils.Exists(unitfile_link); exists {
		err := os.Remove(unitfile_link)
		if err != nil {
			log.StdLogger().Panic(err)
		}
	}

	// write new unitfile
	reader := strings.NewReader(unittext)
	opts, err := unit.Deserialize(reader)
	if err != nil {
		log.StdLogger().Panic(err)
	}

	opts = append(opts, unit.NewUnitOption("Unit", "Description", app.Description))
	opts = append(opts, unit.NewUnitOption("Unit", "Documentation", app.Repository))
	opts = append(opts, unit.NewUnitOption("Service", "ExecStart", filepath.Join(BinPath, app.Name)+" start"))
	opts = append(opts, unit.NewUnitOption("Install", "Alias", fmt.Sprintf("%s.service clustermom-%s.service", app.Name, app.Name)))
	r := unit.Serialize(opts)

	var f *os.File
	exist, _ := utils.Exists(unitfile)
	if exist {
		f, err = os.OpenFile(unitfile, os.O_APPEND, 0666)
		if err != nil {
			log.StdLogger().Panic(err)
		}
	} else {
		f, _ = utils.MakeFile(unitfile)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.StdLogger().Panic(err)
		}
	}(f)
	w := bufio.NewWriter(f)
	_, err = w.ReadFrom(r)
	if err != nil {
		log.StdLogger().Panic(err)
	}
	defer func(w *bufio.Writer) {
		err := w.Flush()
		if err != nil {
			log.StdLogger().Panic(err)
		}
	}(w)

	// link unitfile to /etc/systemd/system/
	if err := os.Link(unitfile, unitfile_link); err != nil {
		log.StdLogger().Panic(err)
	}
}

func ReloadSystemdDaemon() {
	ctx := context.TODO()
	systemd, err := dbus.NewWithContext(ctx)
	if err != nil {
		log.StdLogger().Panic(err)
	}
	defer systemd.Close()
	err = systemd.ReloadContext(ctx)
	if err != nil {
		log.StdLogger().Panic(err)
	}
}
