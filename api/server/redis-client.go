package server

import (
	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
)

type RedisClient redis.Client

func NewRedisClient() *RedisClient {
	// Replaces with your configuration information
	cfg := getConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
		Password: cfg.RedisPasswd,
		DB:       0, // use default DB
	})
	log.Info().Msg("Redis init done ...")

	pong, err := client.Ping().Result()
	if err != nil {
		log.Error().Err(err)
	}
	log.Info().Str("ping", pong).Msg("Redis ping response")
	return (*RedisClient)(client)
}
