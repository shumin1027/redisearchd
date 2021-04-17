package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

/**
 * env			环境dev/test/proc
 * logFile 		日志文件路径
 * level 		日志级别 debug/info/warn/error/dpanic/panic/fatal
 * encoding 	日志编码方式 json/console
 * maxSize 		每个日志文件保存的最大尺寸 单位：M
 * maxBackups 	日志文件最多保存多少个备份
 * maxAge 		文件最多保存多少天
 * compress 	是否压缩
 * app 			应用名称
 */
func NewLogger(env, logFile string, level, encoding string, maxSize int, maxBackups int, maxAge int, compress bool, stdout bool, app string) *zap.Logger {
	core := newCore(logFile, level, encoding, maxSize, maxBackups, maxAge, compress, stdout)
	var logger *zap.Logger
	if env == "proc" {
		logger = zap.New(core, zap.WithCaller(false), zap.Fields(zap.String("app", app)))
	} else {
		logger = zap.New(core, zap.WithCaller(true), zap.Fields(zap.String("app", app)))
	}
	return logger
}

/**
 * zapcore构造
 */
func newCore(logFile string, level, encoding string, maxSize int, maxBackups int, maxAge int, compress bool, stdout bool) zapcore.Core {
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	var l zapcore.Level
	if level != "" {
		l.Set(level)
	}
	atomicLevel.SetLevel(l)

	// 编码器配置
	encoderConfig := zapcore.EncoderConfig{
		NameKey:    "logger",
		MessageKey: "msg",

		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder, // ISO8601 UTC 时间格式

		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,

		CallerKey:    "caller",
		EncodeCaller: zapcore.FullCallerEncoder, // 全路径编码器

		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,

		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志编码
	var enc zapcore.Encoder
	switch encoding {
	case "json":
		enc = zapcore.NewJSONEncoder(encoderConfig)
	case "console":
		enc = zapcore.NewConsoleEncoder(encoderConfig)
	}

	var ws zapcore.WriteSyncer
	if logFile != "" {
		//日志文件路径配置
		hook := lumberjack.Logger{
			Filename:   logFile,    // 日志文件路径
			MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: maxBackups, // 日志文件最多保存多少个备份
			MaxAge:     maxAge,     // 文件最多保存多少天
			Compress:   compress,   // 是否压缩
		}
		if stdout {
			ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
		} else {
			ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook))
		}
	} else {
		ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	}

	return zapcore.NewCore(enc, ws, atomicLevel)
}
