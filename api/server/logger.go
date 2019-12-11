package server

import (
	"fmt"
	"os"

	"github.com/caarlos0/env"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {

	zerolog.TimestampFieldName = "ts"
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = zerolog.New(os.Stdout).With().
		Str("type", "APPLICATION").
		Str("appName", "picturedictionnary-api").
		Str("appVersion", "0.0.1").
		Timestamp().
		Logger()

	cfg := getConfig()
	debug := cfg.IsDebug

	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		output := zerolog.ConsoleWriter{Out: os.Stdout}
		output.FormatMessage = func(i interface{}) string {
			return fmt.Sprintf(" | %s | ", i)
		}
		log.Logger = zerolog.New(os.Stdout).Output(output).With().
			Str("module", "api").
			Timestamp().
			Logger()
	}
}

func getConfig() config {
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing config.go")
	}

	return cfg
}
