package logger

type Logger interface {
	Debug(msg string, data ...map[string]interface{})
	Info(msg string, data ...map[string]interface{})
	Warn(msg string, data ...map[string]interface{})
	Error(msg string, data ...map[string]interface{})
	DPanic(msg string, data ...map[string]interface{})
	Panic(msg string, data ...map[string]interface{})
	Fatal(msg string, data ...map[string]interface{})
}
