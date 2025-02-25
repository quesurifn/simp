package logger

import (
	"github.com/rs/zerolog"
	"os"
)

type Logger struct {
	Instance zerolog.Logger
}

func NewLogger(service string, debug bool) *Logger {
	hostname, _ := os.Hostname()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("role", service).
		Str("host", hostname).
		Logger()

	return &Logger{
		Instance: logger,
	}
}
