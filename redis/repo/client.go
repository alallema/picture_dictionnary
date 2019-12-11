package repo

import (
	"fmt"

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
	fmt.Println("redis init done")
}
