package one

import "github.com/rs/zerolog"

type Config struct {
	ServerPort int
	Debugging  bool

	LogLevel     zerolog.Level
	LogChannel   string
	LogFormatter string
}
