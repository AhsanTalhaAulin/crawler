package conn

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func RedisConnect() error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // Use if required
		DB:       0,  // Choose the appropriate DB
	})

	_, err := client.Ping(context.TODO()).Result()
	if err != nil {
		return err
	}

	RedisClient = client
	log.Println("Connected to Redis")
	return nil
}
