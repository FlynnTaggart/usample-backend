package logger

import (
	"go.uber.org/zap"
)

type ZapAdapter struct {
	ZapLogger *zap.Logger
}

func NewZapAdapter(zapLogger *zap.Logger) *ZapAdapter {
	return &ZapAdapter{ZapLogger: zapLogger}
}

func (a ZapAdapter) Log(lvl string, msg string, data map[string]interface{}) {
	fields := make([]zap.Field, 0, len(data))
	for k, v := range data {
		fields = append(fields, zap.Any(k, v))
	}

	switch lvl {
	case "debug":
		a.ZapLogger.Debug(msg, fields...)
	case "info":
		a.ZapLogger.Info(msg, fields...)
	case "warn":
		a.ZapLogger.Warn(msg, fields...)
	case "error":
		a.ZapLogger.Error(msg, fields...)
	case "fatal":
		a.ZapLogger.Fatal(msg, fields...)
	case "panic":
		a.ZapLogger.Panic(msg, fields...)
	default:
		a.ZapLogger.Error(msg, fields...)
	}
}

func (a ZapAdapter) Debug(msg string, data map[string]interface{}) {
	a.Log("debug", msg, data)
}

func (a ZapAdapter) Info(msg string, data map[string]interface{}) {
	a.Log("info", msg, data)
}

func (a ZapAdapter) Warn(msg string, data map[string]interface{}) {
	a.Log("warn", msg, data)
}

func (a ZapAdapter) Error(msg string, data map[string]interface{}) {
	a.Log("error", msg, data)
}

func (a ZapAdapter) Panic(msg string, data map[string]interface{}) {
	a.Log("panic", msg, data)
}

func (a ZapAdapter) Fatal(msg string, data map[string]interface{}) {
	a.Log("fatal", msg, data)
}
