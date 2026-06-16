package storage

import (
	"context"
	"log"

	"github.com/lanxre/kyokusulib/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Printf("Redis not available at %s — running without cache\n", cfg.RedisAddr)
		return nil
	}

	log.Println("Connected to redis at", cfg.RedisAddr)
	return client
}
