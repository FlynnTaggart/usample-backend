package main

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"io"
	"os"
	"strings"
	"user-service/pkg/logger"
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

	_ = logger.NewLogrusAdapter(initializeLogger(f))
}
