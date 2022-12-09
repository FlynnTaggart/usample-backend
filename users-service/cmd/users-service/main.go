package main

import (
	"context"
	"errors"
	"fmt"
	pgxLogrus "github.com/jackc/pgx-logrus"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	uuid "github.com/vgarvardt/pgx-google-uuid/v5"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
	"user-service/internal/db"
	"user-service/internal/pb"
	"user-service/internal/servers"
	"user-service/internal/services"
	"user-service/pkg/logger"
)

func initializeLogger(logFile *os.File) *logrus.Logger {
	return &logrus.Logger{
		Out:   io.MultiWriter(logFile, os.Stdout),
		Level: logrus.DebugLevel,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}
}

func main() {
	logPath := strings.TrimRight(os.Getenv("LOG_DIR"), "/")
	if _, err := os.Stat(logPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(logPath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	f, err := os.OpenFile(logPath+"/"+os.Getenv("LOG_FILENAME"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed to create logfile %s/%s", logPath, os.Getenv("LOG_FILENAME"))
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("failed to close logfile %s/%s", logPath, os.Getenv("LOG_FILENAME"))
		}
	}(f)

	logrusLogger := initializeLogger(f)

	logrusLoggerAdapted := logger.NewLogrusAdapter(logrusLogger)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		logrusLoggerAdapted.Fatal(fmt.Sprintf("main: failed to open tcp port: %v", err), map[string]interface{}{})
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"))

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal(err)
	}

	tr := &tracelog.TraceLog{
		Logger:   pgxLogrus.NewLogger(logrusLogger),
		LogLevel: tracelog.LogLevelDebug,
	}

	config.ConnConfig.Tracer = tr
	config.MaxConns = 50
	config.MaxConnLifetime = time.Minute * 10
	config.MaxConnIdleTime = time.Minute * 30

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		uuid.Register(conn.TypeMap())
		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		logrusLoggerAdapted.Fatal("failed to connect to pg db", map[string]interface{}{})
	}

	pgxDB := db.NewPgxDB(pool, tr.Logger)

	service := services.NewUsersService(pgxDB, logrusLoggerAdapted)

	server := servers.NewUsersServer(service, logrusLoggerAdapted)

	grpcServer := grpc.NewServer()

	pb.RegisterUsersServiceServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		logrusLogger.Fatal(fmt.Sprintf("main: failed to serve: %s", err.Error()), map[string]interface{}{})
	}
}
