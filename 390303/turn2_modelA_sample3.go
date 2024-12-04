package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go func(ctx context.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered: %v\n", r)
			}
		}()

		select {
		case <-ctx.Done():
			print("Timeout occurred.")
			return
		default:
			file, err := os.Create("example.txt")
			if err != nil {
				fmt.Printf("Error creating file: %v\n", err)
				return
			}
			defer file.Close()

			_, err = file.WriteString("Hello, world!")
			if err != nil {
				fmt.Printf("Error writing to file: %v\n", err)
			}
		}
	}(ctx)
}
