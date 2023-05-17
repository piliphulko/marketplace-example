package logwriter

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitializeLogger(logger *zap.Logger, pathInfoLevel, pathErrorLevel, pathPanicLevel string) error {
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
	logger = zap.New(core) // zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return nil
}

func InitializeLoggerd(pathInfoLevel, pathErrorLevel, pathPanicLevel string) (*zap.Logger, error) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	logFileInfo, err := os.OpenFile(pathInfoLevel, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0755))
	if err != nil {
		return nil, err
	}
	logFileError, err := os.OpenFile(pathErrorLevel, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0755))
	if err != nil {
		return nil, err
	}
	logFilePanic, err := os.OpenFile(pathPanicLevel, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0755))
	if err != nil {
		return nil, err
	}
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(logFileInfo), zapcore.InfoLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(logFileError), zapcore.ErrorLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(logFilePanic), zapcore.PanicLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return logger, nil
}
