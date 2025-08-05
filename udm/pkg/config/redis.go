package config

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var (
	Rdb *redis.Client
	Ctx = context.Background()
)

func ConnectRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "db:6379",
		Password: "", 
		DB:       0, 
		PoolSize:     4000,
		MinIdleConns: 400,  
		PoolTimeout:  30 * time.Second, 
		ConnMaxIdleTime:  5 * time.Minute, 
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}
