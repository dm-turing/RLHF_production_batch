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
		Addr: "localhost:6379",
		DB:   0,
	})

	streamName := "mystream"
	lastID := "0"

	for {
		streams, err := rdb.XRead(ctx, &redis.XReadArgs{
			Streams: []string{streamName, lastID},
			Count:   5, // Read 5 messages at a time
			Block:   0, // Block indefinitely
		}).Result()

		if err != nil {
			log.Fatalf("Failed to read from stream: %v", err)
		}

		for _, stream := range streams {
			for _, message := range stream.Messages {
				fmt.Printf("Stream: %s, ID: %s, Values: %v\n", stream.Stream, message.ID, message.Values)
				lastID = message.ID
			}
		}
	}
}
