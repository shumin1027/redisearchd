package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(conf *LogConf) *zap.Logger {
	core := newCore(conf)
	var logger *zap.Logger
	logger = zap.New(core, zap.WithCaller(conf.WithCaller))
	if conf.AppName != "" {
		logger = logger.WithOptions(zap.Fields(zap.String("app", conf.AppName)))
	}

	return logger
}

/**
 * zapcore构造
 */
func newCore(conf *LogConf) zapcore.Core {
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	var l zapcore.Level
	if conf.Level != "" {
		l.Set(conf.Level)
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

		CallerKey: "caller",
		//EncodeCaller: zapcore.FullCallerEncoder, // 全路径编码器
		EncodeCaller: zapcore.ShortCallerEncoder, // 短路径编码器

		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,

		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志编码
	var enc zapcore.Encoder
	switch conf.Encoding {
	case "json":
		enc = zapcore.NewJSONEncoder(encoderConfig)
	case "console":
		// 仅仅控制台输出的时候，启用着色
		if !conf.Logrotate {
			encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		}
		enc = zapcore.NewConsoleEncoder(encoderConfig)
	}
	var ws zapcore.WriteSyncer

	if conf.Logrotate {
		//日志文件路径配置
		hook := lumberjack.Logger{
			Filename:   conf.Filename,   // 日志文件路径
			MaxSize:    conf.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: conf.MaxBackups, // 日志文件最多保存多少个备份
			MaxAge:     conf.MaxAge,     // 文件最多保存多少天
			Compress:   conf.Compress,   // 是否压缩
			LocalTime:  conf.LocalTime,  // 是否压缩
		}
		if conf.Stdout {
			// 同时输出到 Stdout 和 File
			ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
		} else {
			// 只输出到 File
			ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook))
		}
	} else {
		if conf.Stdout {
			// 只输出到 Stdout
			ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
		}
	}
	return zapcore.NewCore(enc, ws, atomicLevel)
}
