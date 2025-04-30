package redisconfig

import "os"

type config struct {
	HOST string
	PORT string
}

func RedisConfig() *config {

	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("REDIS_PORT")
	if port == "" {
		port = "6379"
	}

	return &config{
		HOST: host,
		PORT: port,
	}
}
