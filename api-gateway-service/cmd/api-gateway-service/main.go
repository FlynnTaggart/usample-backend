package main

import (
	"api-gateway-service/internal/auth"
	"api-gateway-service/internal/auth/handlers"
	"api-gateway-service/internal/users"
	"api-gateway-service/pkg/logger"

	"fmt"
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
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

	app := fiber.New()
	apiRoute := app.Group("/api/v1")

	authClient := auth.InitializeLoginRoute(apiRoute, os.Getenv("AUTH_SVC_URL"), logrusLogger)
	usersClient := users.InitializeUsersRoutes(apiRoute, os.Getenv("USERS_SVC_URL"), logrusLogger, authClient)
	InitializeRegisterRoute(apiRoute, authClient, usersClient)

	if err := app.Listen(os.Getenv("GATEWAY_URL")); err != nil {
		logrusLogger.Fatal(fmt.Sprintf("Server is not running! Reason: %v", err), map[string]interface{}{})
	}
}

func InitializeRegisterRoute(a fiber.Router, authClient *auth.ServiceClient, usersClient *users.ServiceClient) {
	a.Post("/register", func(ctx *fiber.Ctx) error {
		return handlers.Register(ctx, authClient.Client, usersClient.Client)
	})
}
