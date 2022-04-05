package logger

import (
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
}

type LoggerImpl struct {
	*log.Entry
}

func NewLogger(c echo.Context) Logger {
	return &LoggerImpl{
		makeLogEntry(c),
	}
}

func makeLogEntry(c echo.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}
	start := c.Get("start").(time.Time)
	return log.WithFields(log.Fields{
		"at":         time.Now().Format("2006-01-02 15:04:05"),
		"method":     c.Request().Method,
		"status":     c.Response().Status,
		"uri":        c.Request().URL.String(),
		"ip":         c.Request().RemoteAddr,
		"latency_ns": time.Since(start).String(),
	})
}
