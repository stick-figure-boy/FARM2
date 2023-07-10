package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

var logger *zerolog.Logger

const defaultSkip = 4

func init() {
	zerolog.TimeFieldFormat = "2006-01-02T15:04:-5.999Z07:00"
	host, _ := os.Hostname()
	log := zerolog.New(os.Stderr).With().Timestamp().Str("host", host).Stack().Logger()
	logger = &log
}

func Info(val ...any) {
	var msg []string
	for _, v := range val {
		msg = append(msg, fmt.Sprintf("%+v", v))
	}
	logger.Info().Msg(strings.Join(msg, ", "))
}

func Warn(val ...any) {
	val = append(val, getStackTrace(defaultSkip))

	var msg []string
	for _, v := range val {
		msg = append(msg, fmt.Sprintf("%+v", v))
	}
	logger.Warn().Msg(strings.Join(msg, ", "))
}

func Err(val ...any) {
	val = append(val, getStackTrace(defaultSkip))

	var errs []error
	for _, v := range val {
		errs = append(errs, errors.New(fmt.Sprintf("%+v", v)))
	}
	logger.Error().Errs("message", errs).Msg("")
}

func getStackTrace(skip int) string {
	defer recover()

	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}

	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	if n == 0 {
		return ""
	}

	pc = pc[:n]
	frames := runtime.CallersFrames(pc)
	var funcs []string
	for {
		frame, more := frames.Next()
		if !more {
			break
		}

		f := frame.File
		if strings.Contains(f, "runtime/") || strings.HasSuffix(f, "extern.go") || strings.HasSuffix(f, "logger.go") || strings.HasSuffix(f, "signal_unix.go") {
			continue
		}
		fmt.Println(frame.Function)
		funcs = append(funcs, fmt.Sprintf("%s, func: %s, line: %d", f, frame.Function, frame.Line))
	}
	for i, j := 0, len(funcs)-1; i < j; i, j = i+1, j-1 {
		funcs[i], funcs[j] = funcs[j], funcs[i]
	}

	return fmt.Sprintf("%s, line: %d, caller: %s", file, line, strings.Join(funcs, " >> "))
}
