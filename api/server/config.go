package server

type config struct {
	IsDebug     bool   `env:"DEBUG_ENVIRONMENT" envDefault:"false"`
	RedisHost   string `env:"REDIS_HOST" envDefault:"localhost"`
	RedisPort   string `env:"REDIS_PORT" envDefault:"6379"`
	RedisPasswd string `env:"REDIS_PASSWORD" envDefault:""`
	ApiPort     string `env:"API_PORT" envDefault:"4000"`
}
