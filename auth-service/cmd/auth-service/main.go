package main

import (
	"auth-service/internal/db"
	"auth-service/internal/pb"
	"auth-service/internal/servers"
	"auth-service/internal/services"
	"auth-service/pkg/logger"
	"auth-service/utils"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"net"
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
	zapLogger := logger.NewZapAdapter(initializeLogger())

	DB := db.NewRedisDB(os.Getenv("REDIS_URL"), os.Getenv("REDIS_PASSWORD"))

	jwt := utils.NewJwtWrapper(os.Getenv("JWT_SECRET_KEY"), "auth-service", 1)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		zapLogger.Fatal(fmt.Sprintf("main: failed to open tcp port: %v", err))
	}

	service := services.NewAuthService(*DB, *jwt)

	server := servers.NewAuthServer(service)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		zapLogger.Fatal(fmt.Sprintf("main: failed to serve: %s", err.Error()))
	}
}
