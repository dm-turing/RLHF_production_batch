package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// Create a Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Use your Redis server address
		DB:   0,                // Use default DB
	})

	// Write a message to the 'mystream' stream
	streamName := "mystream"
	_, err := rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: streamName,
		Values: map[string]interface{}{"field1": "value1"},
	}).Result()

	if err != nil {
		log.Fatalf("Failed to write to stream: %v", err)
	}

	fmt.Println("Message published to stream")
}
