// Package logging provides structured logging with log.
package logging

import (
	"github.com/changyoungkwon/gxample/internal/config"
	log "github.com/sirupsen/logrus"
)

// Logger for global, initialize by init
var Logger *log.Logger

func init() {
	Logger = log.New()
	Logger.SetFormatter(&log.TextFormatter{
		DisableTimestamp: false,
	})
	level := config.Get().Log.Level
	if level == "" {
		level = "error"
	}
	l, err := log.ParseLevel(level)
	if err != nil {
		log.Fatalf("invalid log-level %s, set default to error", level)
		l, _ = log.ParseLevel("error")
	}
	Logger.SetLevel(l)
}
