// Package logging provides structured logging with log.
package logging

import (
	"github.com/changyoungkwon/gxample/internal/config"
	log "github.com/sirupsen/logrus"
)

// Logger for global, initialize by init
var Logger *log.Logger

// Fatalf wraps logger's Fatalf
func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args...)
}

// Errorf wraps logger's Errorf
func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args...)
}

// Warnf wraps logger's Warnf
func Warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args...)
}

// Infof wraps logger's Infof
func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}

// Debugf wraps logger's Debugf
func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}

// Fatal wraps logger's Fatal
func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

// Error wraps logger's Error
func Error(args ...interface{}) {
	Logger.Error(args...)
}

// Warn wraps logger's Warn
func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

// Info wraps logger's Info
func Info(args ...interface{}) {
	Logger.Info(args...)
}

// Debug wraps logger's Debug
func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

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
