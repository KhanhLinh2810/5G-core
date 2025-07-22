package config

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var (
	Rdb *redis.Client
	Ctx = context.Background()
)

func ConnectRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", 
		DB:       0,  
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}
