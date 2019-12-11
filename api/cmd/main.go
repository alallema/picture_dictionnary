package main

import (
	"fmt"

	"github.com/alallema/picture_dictionnary.git/api/server"
	"github.com/rs/zerolog/log"
)

func errorHandling() {
	if err := recover(); err != nil {
		err2 := fmt.Errorf("recover error not nil: %g", err) // format err from interface type to error type
		log.Error().Err(err2).Msg("recovered from panic")
	}
}

func main() {
	defer errorHandling()
	server.InitLogger()
	log.Info().Msg("Launching picture dictionnary api ...")
	server := server.NewServer()
	server.Run()
}
