package main

import (
	"api-gateway-service/internal/auth"
	"api-gateway-service/pkg/logger"
	
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	zapLogger := logger.NewZapAdapter(initializeLogger())
	defer func() {
		err := zapLogger.ZapLogger.Sync()
		if err != nil {
			log.Fatalf("zap: %v", err)
		}
	}()

	app := fiber.New()

	_ = auth.InitializeRoutes(app.Group("/api"), os.Getenv("AUTH_SVC_URL"), zapLogger)

	if err := app.Listen(os.Getenv("GATEWAY_URL")); err != nil {
		zapLogger.Fatal(fmt.Sprintf("Server is not running! Reason: %v", err))
	}
}
