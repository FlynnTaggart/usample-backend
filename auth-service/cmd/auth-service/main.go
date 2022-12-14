package main

import (
	"auth-service/internal/db"
	"auth-service/internal/pb"
	"auth-service/internal/servers"
	"auth-service/internal/services"
	"auth-service/pkg/logger"
	"auth-service/utils"
	"errors"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"io"
	"strings"

	"fmt"
	"net"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

func initializeLogger(logFile *os.File) *logrus.Logger {
	log := &logrus.Logger{
		Out:   io.MultiWriter(logFile, os.Stdout),
		Level: logrus.DebugLevel,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}
	return log
}

func main() {
	logPath := strings.TrimRight(os.Getenv("LOG_DIR"), "/")
	if _, err := os.Stat(logPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(logPath, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}
	f, err := os.OpenFile(logPath+"/"+os.Getenv("LOG_FILENAME"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile " + logPath + "/" + os.Getenv("LOG_FILENAME"))
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("Failed to close logfile " + logPath + "/" + os.Getenv("LOG_FILENAME"))
		}
	}(f)

	logrusLogger := logger.NewLogrusAdapter(initializeLogger(f))

	DB := db.NewRedisDB(os.Getenv("REDIS_URL"), os.Getenv("REDIS_PASSWORD"))

	jwt := utils.NewJwtWrapper(os.Getenv("JWT_SECRET_KEY"), "auth-service", 1)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		logrusLogger.Fatal(fmt.Sprintf("main: failed to open tcp port: %v", err), map[string]interface{}{})
	}

	service := services.NewAuthService(*DB, *jwt)

	server := servers.NewAuthServer(service)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		logrusLogger.Fatal(fmt.Sprintf("main: failed to serve: %s", err.Error()), map[string]interface{}{})
	}
}
