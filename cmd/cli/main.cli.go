package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// 1736151196.270052 -> 2025-01-06T15:13:16.268+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// timestap (ts) -> Time
	encodeConfig.TimeKey = "time"

	// from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// "caller":"cli/main.cli.go:30"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

func GetWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_RDWR, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)

}

func main() {
	// sugar := zap.NewExample().Sugar()

	// sugar.Info("Hello name : %s, age : %d", "hehehe", 123) // like ftm.Frintf()

	// loger

	// loger := zap.NewExample()
	// loger.Info("Hello")

	// New Example
	// NewLogger := zap.NewExample()
	// NewLogger.Info("Hello NewExample")
	// // {"level":"info","msg":"Hello NewExample"}
	// fmt.Println("------------------------------")

	// // Development
	// NewLogger, _ = zap.NewDevelopment()
	// NewLogger.Info("Hello NewDevelopment")
	// // 2025-01-06T15:13:16.268+0700    INFO    cli/main.cli.go:26      Hello NewDevelopment
	// fmt.Println("------------------------------")

	// // Production
	// NewLogger, _ = zap.NewProduction()
	// NewLogger.Info("Hello NewDevelopment")
	// // {"level":"info","ts":1736151196.270052,"caller":"cli/main.cli.go:31","msg":"Hello NewDevelopment"}
	// fmt.Println("------------------------------")

	encoder := GetEncoderLog()
	sync := GetWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())
	logger.Info("Info log", zap.Int("line", 1))
	fmt.Println("-------------------")
	logger.Error("Error log", zap.Int("line", 2))
}
