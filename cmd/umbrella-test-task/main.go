package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"middleware/internal/pkg/app"
)

func main() {
	// Установка формата времени в логах
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	a, err := app.New()
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	err = a.Run()
	if err != nil {
		log.Error().Err(err).Msg("")
	}
}
