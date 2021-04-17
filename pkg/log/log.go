package log

import (
	"gitlab.xtc.home/xtc/redisearchd/app"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var zapLogger *zap.Logger
var stdLogger *log.Logger

func init() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig = zapcore.EncoderConfig{
		NameKey:    "logger",
		MessageKey: "msg",

		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder, // ISO8601 UTC 时间格式

		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalColorLevelEncoder, // 大写彩色 INFO

		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder, // 全路径编码器

		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,

		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeName:     zapcore.FullNameEncoder,
	}
	cfg.Encoding = "console"
	var err error
	zapLogger, err = cfg.Build(zap.AddCaller())
	if err != nil {
		log.Panic(err)
	}
	stdLogger = zap.NewStdLog(zapLogger)
}

func Init(env, level, encoding, filePath string, stdout bool) {
	zapLogger = NewLogger(env, filePath, level, encoding, 128, 30, 7, true, stdout, app.Name)
	Logger().Info("zap logger init success")
}

func Logger() *zap.Logger {
	return zapLogger
}

func StdLogger() *log.Logger {
	return stdLogger
}

func Sugar() *zap.SugaredLogger {
	return zapLogger.Sugar()
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(msg string, fields ...zap.Field) {
	Logger().Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(msg string, fields ...zap.Field) {
	Logger().Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(msg string, fields ...zap.Field) {
	Logger().Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(msg string, fields ...zap.Field) {
	Logger().Error(msg, fields...)
}

// DPanic logs a message at DPanicLevel. The message includes any fields
// passed at the log site, as well as any fields accumulated on the logger.
//
// If the logger is in development mode, it then panics (DPanic means
// "development panic"). This is useful for catching errors that are
// recoverable, but shouldn't ever happen.
func DPanic(msg string, fields ...zap.Field) {
	Logger().DPanic(msg, fields...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func Panic(msg string, fields ...zap.Field) {
	Logger().Panic(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func Fatal(msg string, fields ...zap.Field) {
	Logger().Fatal(msg, fields...)
}
