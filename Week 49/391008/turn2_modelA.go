package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	c.AddFunc("0 2 * * *", func() {
		fmt.Println("Running backup job...")
		// Call your backup function here
	})
	c.Start()
	defer c.Stop()

	select {} // Keep the program running
}
