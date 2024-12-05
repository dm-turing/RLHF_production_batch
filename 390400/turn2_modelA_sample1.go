package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Create a Redis client with default pooling options
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Perform a Redis operation
	val, err := client.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println("Error getting key:", err)
		return
	}
	fmt.Println("Value from Redis:", val)
}
