package gofnd

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

const (
	logsDir = "logs"
)

// TODO: upd logging: source: https://www.honeybadger.io/blog/golang-logging/

func SetupLogger(fileName string, writeToFile bool) {
	log.SetFormatter(&log.JSONFormatter{})

	logLevel := log.TraceLevel

	logLevelRaw := "TRACE"

	switch logLevelRaw {
	case "INFO":
		logLevel = log.InfoLevel
	case "DEBUG":
		logLevel = log.DebugLevel
	case "TRACE":
		logLevel = log.TraceLevel
	}

	log.SetLevel(logLevel)

	if writeToFile {
		if err := os.MkdirAll(logsDir, 0777); err != nil {
			log.Fatalf("gofnd[.SetupLogger][1]: %v", err)
		}

		logFilePath := filepath.Join(logsDir, fmt.Sprintf("%v.log", fileName))

		file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			log.Fatalf("gofnd[.SetupLogger][2]: %v", err)
		}

		// TODO: need close file???
		log.SetOutput(file)
	}
}
