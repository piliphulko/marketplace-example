package logwriter

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapgrpc"
)

type logSync func()

func InitZapLog(logger **zap.Logger, pathLogFile string, level zapcore.Level) (logSync, error) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	logFile, err := os.OpenFile(pathLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0755))
	if err != nil {
		return nil, err
	}
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(logFile), level),
	)
	createdLogger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(level))
	*logger = createdLogger
	return func() { createdLogger.Sync() }, nil
}

func InitZapLogStdout(logger **zap.Logger, level zapcore.Level) logSync {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(os.Stdout), level),
	)
	createdLogger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(level))
	*logger = createdLogger
	return func() { createdLogger.Sync() }
}

func InitZapLogGRPC(logger **zapgrpc.Logger, pathLogFile string, level zapcore.Level) (logSync, error) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	logFile, err := os.OpenFile(pathLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0755))
	if err != nil {
		return nil, err
	}
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(logFile), level),
	)
	createdLogger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(level))
	*logger = zapgrpc.NewLogger(createdLogger)
	return func() { createdLogger.Sync() }, nil
}

func NewZapLogStdoutGRPC(logger **zapgrpc.Logger, level zapcore.Level) logSync {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(os.Stdout), level),
	)
	createdLogger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(level))
	*logger = zapgrpc.NewLogger(createdLogger)
	return func() { createdLogger.Sync() }
}

/*
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
*/
