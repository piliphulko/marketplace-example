package logwriter

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapgrpc"
)

type logSync func()

func InitializeLogger(logger **zap.Logger, pathInfoLevel, pathErrorLevel, pathPanicLevel string) error {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	logFileInfo, err := os.OpenFile(pathInfoLevel, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0755))
	if err != nil {
		return err
	}
	logFileError, err := os.OpenFile(pathErrorLevel, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0755))
	if err != nil {
		return err
	}
	logFilePanic, err := os.OpenFile(pathPanicLevel, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0755))
	if err != nil {
		return err
	}
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(logFileInfo), zapcore.InfoLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(logFileError), zapcore.ErrorLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(logFilePanic), zapcore.PanicLevel),
	)
	*logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return nil
}

func InitializeLoggerGRPC(logger **zapgrpc.Logger, pathLogger string) (error, logSync) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	logFile, err := os.OpenFile(pathLogger, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0755))
	if err != nil {
		return err, nil
	}
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(logFile), zapcore.InfoLevel),
	)
	basicLogger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	*logger = zapgrpc.NewLogger(basicLogger)
	return nil, func() { basicLogger.Sync() }
}

func InitStdoutLoggerGRPC(logger **zapgrpc.Logger) logSync {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
	)
	basicLogger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	*logger = zapgrpc.NewLogger(basicLogger)
	return func() { basicLogger.Sync() }
}
