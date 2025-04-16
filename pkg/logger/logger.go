package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.Logger

func InitWithRotation(logPath string, logLevel string) {
	// Parse log level
	level := getLogLevel(logLevel)

	// Configure lumberjack for log rotation
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logPath, // Log file path
		MaxSize:    500,     // Max size in megabytes before rotation
		MaxBackups: 5,       // Maximum number of old log files to retain
		MaxAge:     30,      // Maximum days to retain old log files
		Compress:   true,    // Compress rotated files with gzip
	})

	// Configure encoder
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "timestamp",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Create the core with file and console output
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)

	// Also log to console for immediate visibility
	consoleOutput := zapcore.Lock(os.Stdout)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, w, level),
		zapcore.NewCore(consoleEncoder, consoleOutput, level),
	)

	// Create the logger
	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	Log.Info("Logger initialized with rotation",
		zap.String("log_path", logPath),
		zap.String("log_level", logLevel))
}

func Init(logLevel string) {
	level := getLogLevel(logLevel)

	config := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	var err error
	Log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func getLogLevel(logLevel string) zapcore.Level {
	switch logLevel {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func Info(message string, fields ...zap.Field) {
	Log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	Log.Debug(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	Log.Fatal(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	Log.Warn(message, fields...)
}

func Sync() error {
	return Log.Sync()
}
