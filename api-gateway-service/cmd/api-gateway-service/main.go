package main

import (
	"api-gateway-service/internal/auth"
	"api-gateway-service/pkg/zapadapter"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

func initializeLogger() *zap.Logger {
	config := zap.NewDevelopmentEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	logFile, _ := os.OpenFile("./logs/server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func main() {
	logger := zapadapter.ZapAdapter{ZapLogger: initializeLogger()}
	defer func() {
		err := logger.ZapLogger.Sync()
		if err != nil {
			log.Fatalf("zap: %v", err)
		}
	}()

	app := fiber.New()

	authSvc := auth.InitializeRoutes(app, os.Getenv("AUTH_SVC_URL"), logger)
}
