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

func (a ZapAdapter) Log(lvl string, msg string, data ...map[string]interface{}) {
	var fields []zap.Field
	for _, d := range data {
		for k, v := range d {
			field := zap.Any(k, v)
			fields = append(fields, field)
		}
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
	case "dpanic":
		a.ZapLogger.DPanic(msg, fields...)
	default:
		a.ZapLogger.Error(msg, fields...)
	}
}

func (a ZapAdapter) Debug(msg string, data ...map[string]interface{}) {
	a.Log("debug", msg, data...)
}

func (a ZapAdapter) Info(msg string, data ...map[string]interface{}) {
	a.Log("info", msg, data...)
}

func (a ZapAdapter) Warn(msg string, data ...map[string]interface{}) {
	a.Log("warn", msg, data...)
}

func (a ZapAdapter) Error(msg string, data ...map[string]interface{}) {
	a.Log("error", msg, data...)
}

func (a ZapAdapter) DPanic(msg string, data ...map[string]interface{}) {
	a.Log("dpanic", msg, data...)
}

func (a ZapAdapter) Panic(msg string, data ...map[string]interface{}) {
	a.Log("panic", msg, data...)
}

func (a ZapAdapter) Fatal(msg string, data ...map[string]interface{}) {
	a.Log("fatal", msg, data...)
}
