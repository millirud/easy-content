package logger

import (
	"io"
	"log"

	"github.com/rs/zerolog"
)

func New(w io.Writer, level string) *zerolog.Logger {

	zlevel, err := zerolog.ParseLevel(level)

	if err != nil {
		log.Fatal(err)
	}

	logger := zerolog.New(w).
		With().
		Timestamp().
		Logger().
		Level(zlevel)

	return &logger
}
