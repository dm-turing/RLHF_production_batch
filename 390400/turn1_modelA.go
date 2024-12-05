package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Set up a Redis client
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Write to a Redis Stream
	streamName := "myStream"
	msgID, err := client.XAdd(ctx, &redis.XAddArgs{
		Stream: streamName,
		MaxLen: int64(100), // Limit to 100 messages
		Values: map[string]string{"field": "value"},
	}).Result()

	if err != nil {
		fmt.Println("Error writing to Redis Stream:", err)
		return
	}
	fmt.Println("Message added to stream:", streamName, "with ID:", msgID)

	// Read from the Redis Stream
	streamConsumer := "consumer1"
	group := "group1"

	// Block on new messages with XReadGroup
	stream, err := client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    group,
		Consumer: streamConsumer,
		Streams:  []string{"$"}, // Start from the last consumed message
		Count:    int64(1),      // Read one message at a time
	}).Result()

	if err != nil {
		fmt.Println("Error reading from Redis Stream:", err)
		return
	}

	if len(stream[0].Messages) > 0 {
		msg := stream[0].Messages[0]
		fmt.Println("Read message from stream:", streamName, "ID:", msg.ID, "Fields:", msg.Values)

		// Acknowledge the message
		err := client.XAck(ctx, streamName, group, streamConsumer, msg.ID).Err()
		if err != nil {
			fmt.Println("Error acknowledging message:", err)
			return
		}
		fmt.Println("Message acknowledged successfully.")
	} else {
		fmt.Println("No new messages available.")
	}
}
