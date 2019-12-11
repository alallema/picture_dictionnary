package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type Server struct {
	Client *RedisClient
	Router *mux.Router
}

func NewServer() *Server {

	s := &Server{
		Client: NewRedisClient(),
	}
	s.Router = NewRouter(s)
	s.Router.Use(headerMiddleware)
	return s

}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func (server *Server) Run() {
	cfg := getConfig()
	log.Info().Str("port", cfg.ApiPort).Msgf("Starting server")
	err := http.ListenAndServe(":"+cfg.ApiPort, server.Router)
	if err != nil {
		log.Error().Err(err)
	}
}
