package logger

import (
	"errors"
	config2 "github.com/ArtisanCloud/go-framework/config"
	UBT "github.com/ArtisanCloud/ubt-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var UBTHandler *UBT.UBT

func SetupLog() error {

	UBTHandler = UBT.Init(config2.UBTConfig)
	//fmt.Dump(config2.UBTConfig)
	if UBTHandler == nil {
		return errors.New("init ubt error")
	}
	return nil
}

func init() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	Logger, _ = config.Build()

	//Logger.Info("123", zap.String("key", "value"))

}

func Info(msg string, fields map[string]interface{}) {

	zapFields := map2ZapFields(fields)

	Logger.Info(msg, zapFields...)

}

func Warn(msg string, fields map[string]interface{}) {

	zapFields := map2ZapFields(fields)

	Logger.Warn(msg, zapFields...)

}

func Error(msg string, fields map[string]interface{}) {

	zapFields := map2ZapFields(fields)

	Logger.Error(msg, zapFields...)

}

func Fatal(msg string, fields map[string]interface{}) {

	zapFields := map2ZapFields(fields)

	Logger.Fatal(msg, zapFields...)

}

func Debug(msg string, fields map[string]interface{}) {

	zapFields := map2ZapFields(fields)

	Logger.Debug(msg, zapFields...)

}

func Panic(msg string, fields map[string]interface{}) {

	zapFields := map2ZapFields(fields)

	Logger.Panic(msg, zapFields...)

}

func map2ZapFields(m map[string]interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(m))
	for k, v := range m {
		fields = append(fields, zap.Any(k, v))
	}
	return fields
}
