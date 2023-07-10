package logger_test

import (
	"errors"
	"testing"

	"github.com/hiroki-Fukumoto/farm2/logger"
)

func TestInfo(t *testing.T) {
	logger.Info("info log")
	logger.Info("info log", "info log2")
}

func TestWarn(t *testing.T) {
	logger.Warn("warning log")
}

func TestErr(t *testing.T) {
	errFunc1()

	e := struct {
		message string
		reason  string
	}{
		message: "ERROR",
		reason:  "ERROR DETAIL",
	}
	logger.Err(e)
}

func errFunc1() {
	logger.Err(errors.New("error log"))
}
