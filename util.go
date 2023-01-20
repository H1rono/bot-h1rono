package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func SetupLogging() {
	// TRACE, DEBUG, INFO, WARNING, ERROR, FATAL, PANIC
	l := os.Getenv("LOG_LEVEL")
	switch l {
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO", "":
		// l = "" の場合に後の出力が残念になるため
		l = "INFO"
		log.SetLevel(log.InfoLevel)
	case "WARNING":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	case "PANIC":
		log.SetLevel(log.PanicLevel)
	default:
		log.Fatalf("Unexpected environment variable LOG_LEVEL=%s", l)
	}
	log.Infof("log level is set at %s", l)
}
