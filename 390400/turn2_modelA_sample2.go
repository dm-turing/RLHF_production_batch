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
	pipeline.Set(ctx, "key1", "value1", 0)
	pipeline.Set(ctx, "key2", "value2", 0)
	pipeline.Get(ctx, "key1")
	pipeline.Get(ctx, "key2")

	// Execute the pipeline and retrieve results
	_, err := pipeline.Exec(ctx)
	if err != nil {
		fmt.Println("Error executing pipeline:", err)
		return
	}

	// Results are returned in the same order as the commands
	fmt.Println("Key1 value:", pipeline.Result(1))
	fmt.Println("Key2 value:", pipeline.Result(2))
}
