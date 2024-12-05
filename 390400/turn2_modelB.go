package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",               // no password set
		DB:           0,                // use default DB
		PoolSize:     10,               // maximum number of socket connections
		MinIdleConns: 5,                // initial number of idle connections
		PoolTimeout:  30 * time.Second, // time to wait for a connection before timing out
	})

	ctx := context.Background()

	err := client.Watch(ctx, func(tx *redis.Tx) error {
		_, err := tx.Pipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, "key1", "value1", 0)
			pipe.Set(ctx, "key2", "value2", 0)
			pipe.Incr(ctx, "counter")
			return nil
		})
		return err
	}, "key1", "key2")

	if err != nil {
		fmt.Println("Error executing batch:", err)
	}

	pipe := client.Pipeline()
	pipe.Set(ctx, "key1", "value1", 0)
	pipe.Set(ctx, "key2", "value2", 0)
	incr := pipe.Incr(ctx, "counter")

	_, err = pipe.Exec(ctx)
	if err != nil {
		fmt.Println("Pipeline error:", err)
	}

	fmt.Println("Incremented counter:", incr.Val())
}
