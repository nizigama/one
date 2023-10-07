package one

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"io"
	"os"
)

type Logger struct {
	*zerolog.Logger
}

const (
	formatterJson   = "json"
	formatterPretty = "pretty"
	channelStdOut   = "stdOut"
)

func NewLogger(level zerolog.Level, formatter string, channel string) *Logger {

	zerolog.SetGlobalLevel(level)
	zerolog.TimeFieldFormat = "2006/01/02 15:04:05"
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	var output io.Writer
	var destination io.Writer

	switch channel {
	case channelStdOut:
		destination = os.Stdout
	default:
		destination = os.Stdout
	}

	switch formatter {
	case formatterPretty:
		output = zerolog.ConsoleWriter{Out: destination, TimeFormat: "2006/01/02 15:04:05"}
	default:
		output = destination
	}

	zeroLogger := zerolog.New(output).With().Timestamp().Caller().Logger()

	return &Logger{&zeroLogger}
}
