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
		file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err) 
		}

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.AddSync(file),                                    
			zap.InfoLevel,                                            
		)

		Logger = zap.New(core)
	})
}

func Log(level, nameFunc, event string, err error, additionalParams ...interface{}) {
	InitLogger()

	eventPrefix := "Event: "

	switch level {
	case "Info":
		Logger.Info(eventPrefix + event)
	case "Error":
		if err != nil {
			Logger.Error(
				err.Error(),
				zap.String("event", event),
				zap.String("func", nameFunc),
				zap.Any("param", additionalParams),
			)
		} else {
			Logger.Error(
				eventPrefix+event,
				zap.String("func", nameFunc),
				zap.Any("param", additionalParams),
			)
		}
	case "Warning":
		Logger.Warn(
			eventPrefix+event,
			zap.String("func", nameFunc),
			zap.Any("param", additionalParams),
		)
	}
}