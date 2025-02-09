package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

var (
	Logger *zap.Logger
	once   sync.Once
)

func InitLogger() {
	once.Do(func() {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		// encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			zap.InfoLevel,
		)

		Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	})
}

func Log(level, event string, err error, additionalParams ...interface{}) {
	if Logger == nil {
		InitLogger()
	}

	fields := make([]zapcore.Field, 0, len(additionalParams)+1)
	for _, param := range additionalParams {
		fields = append(fields, zap.Any("param", param))
	}

	fields = append(fields, zap.String("event", event))

	switch level {
	case "Debug":
		Logger.Debug(event, fields...)
	case "Info":
		Logger.Info(event, fields...)
	case "Warning":
		Logger.Warn(event, fields...)
	case "Error":
		if err != nil {
			fields = append(fields, zap.Error(err))
		}
		Logger.Error(event, fields...)
	default:
		fields = append(fields, zap.String("original_event", event), zap.String("log_level", level))
		Logger.Info("Unknown log level", fields...)
	}
}
