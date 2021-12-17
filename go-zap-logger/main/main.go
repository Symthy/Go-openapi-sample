package main

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	//runLogLevelMultiFileBaseLogger()
	runIndividualMultiLogger()
	// zapLogger, _ := zap.NewProduction()
	// zapGormLogger := zapgorm2.New(zap.L())
}

func runIndividualMultiLogger() {
	infoFile, _ := os.OpenFile("./log/info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	errorFile, _ := os.OpenFile("./log/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	infoCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // フォーマット：コンソール出力形式
		zapcore.AddSync(infoFile),
		zapcore.InfoLevel, // 出力対象：INFO以上
	)

	errorCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),             // フォーマット：json形式
		zapcore.NewMultiWriteSyncer(os.Stderr, errorFile), // コンソール出力とファイル出力
		zapcore.ErrorLevel,                                // 出力対象： ERROR以上
	)

	logger1 := zap.New(infoCore)
	logger2 := zap.New(errorCore)

	logger1.Warn("test warn by logger1")
	logger2.Error("test error by logger2")
}

func runLogLevelMultiFileBaseLogger() {
	infoFile, _ := os.OpenFile("./log/info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	errorFile, _ := os.OpenFile("./log/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	infoCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // フォーマット：コンソール出力形式
		zapcore.AddSync(infoFile),
		zapcore.InfoLevel, // 出力対象：INFO以上
	)

	errorCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),             // フォーマット：json形式
		zapcore.NewMultiWriteSyncer(os.Stderr, errorFile), // コンソール出力とファイル出力
		zapcore.ErrorLevel,                                // 出力対象： ERROR以上
	)

	opts := []zap.Option{}
	opts = append(opts, zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewSamplerWithOptions(
			core,
			time.Second,
			100, //Initial
			100, //Thereafter,
		)
	}))
	logger := zap.New(zapcore.NewTee(infoCore, errorCore), opts...)

	logger.Debug("test debug")
	// 2021-12-17T21:43:41.109+0900	info	test info	{"field": "value"}
	logger.Info("test info", zap.String("field", "value"))
	// 2021-12-17T17:40:21.160+0900	warn	test warn
	logger.Warn("test warn")
	// 2021-12-17T21:43:41.156+0900	error	test error
	logger.Error("test error")
}
