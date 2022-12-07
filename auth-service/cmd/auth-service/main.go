package main

import (
	"auth-service/internal/db"
	"auth-service/internal/pb"
	"auth-service/internal/servers"
	"auth-service/internal/services"
	"auth-service/pkg/logger"
	"auth-service/utils"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"io"

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
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%\n",
		},
	}
	return log
}

func main() {
	f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + "log.txt")
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

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
