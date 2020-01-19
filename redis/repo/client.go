package repo

import (
	"log"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func init() {
	// Replaces with your configuration information
	Client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Active for local test
		// Addr: "redis:6379", // Active for docker
		// Addr: "url:6379", // Active to prod without docker
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	log.Printf("redis init done")

	pong, err := Client.Ping().Result()
	log.Println(pong, err)
	if err != nil {
		log.Printf("can't connect")
	}
}
