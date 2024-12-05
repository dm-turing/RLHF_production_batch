package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Create a Redis client
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Set up a pipeline
	pipeline := client.Pipeline()

	// Add multiple commands to the pipeline
	for i := 0; i < 1000; i++ {
		pipeline.Set(ctx, fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i), 0)
	}

	// Send all commands
	_, err := pipeline.Exec(ctx)
	if err != nil {
		fmt.Println("Error executing pipeline:", err)
		return
	}

	fmt.Println("Pipelined commands executed successfully.")
}
