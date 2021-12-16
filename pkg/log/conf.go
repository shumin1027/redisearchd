package log

import (
	"os"
	"path/filepath"
)

const logPath = "/usr/local/clustermom/var/log/"

// LogConf
/** 日志配置
 * Level 		日志级别 debug/info/warn/error/dpanic/panic/fatal
 * Encoding 	日志编码方式 json/console
 * AppName 		应用名称
 * Stdout		控制台输出
 * Logrotate	日志轮转
 */
type LogConf struct {
	AppName    string `json:"appname" yaml:"appname"`
	Level      string `json:"level" yaml:"level"`
	Encoding   string `json:"encoding" yaml:"encoding"`
	Stdout     bool   `json:"stdout" yaml:"stdout"`
	Logrotate  bool   `json:"logrotate" yaml:"logrotate"`
	WithCaller bool   `json:"withcaller" yaml:"withcaller"`
	LogRotateConf
}

// LogRotateConf
/** 日志轮转配置
 * LogFile		日志文件路径
 * MaxSize 		每个日志文件保存的最大尺寸 单位：M
 * MaxBackups 	日志文件最多保存多少个备份
 * MaxAge 		文件最多保存多少天
 * Compress 	是否压缩
 */
type LogRotateConf struct {
	// Filename is the file to write logs to.  Backup log files will be retained
	// in the same directory.  It uses <processname>-lumberjack.log in
	// os.TempDir() if empty.
	Filename string `json:"filename" yaml:"filename"`

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `json:"maxsize" yaml:"maxsize"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `json:"maxage" yaml:"maxage"`

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int `json:"maxbackups" yaml:"maxbackups"`

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTime bool `json:"localtime" yaml:"localtime"`

	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool `json:"compress" yaml:"compress"`
}

// DevLogConf 开发环境默认配置
func DevLogConf() *LogConf {
	conf := &LogConf{
		Level:      "debug",
		Encoding:   "console",
		WithCaller: true,
		Stdout:     true,
		Logrotate:  false,
	}
	return conf
}

// ProcLogConf 生产环境默认配置
func ProcLogConf() *LogConf {
	rotate := LogRotateConf{
		Filename: func() string {
			exe, _ := os.Executable()
			_, app := filepath.Split(exe)
			logfile := filepath.Join(logPath, app, app+".log")
			return logfile
		}(),
		MaxSize:    128,
		MaxAge:     7,
		MaxBackups: 30,
		LocalTime:  true,
		Compress:   true,
	}
	conf := &LogConf{
		AppName: func() string {
			exe, _ := os.Executable()
			_, app := filepath.Split(exe)
			return app
		}(),
		Level:         "info",
		Encoding:      "json",
		WithCaller:    false,
		Stdout:        true,
		Logrotate:     true,
		LogRotateConf: rotate,
	}
	return conf
}
