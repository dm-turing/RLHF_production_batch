package main

import (
	"fmt"
	"net/http"
	"sync"
)

func fetch(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Fetched:", url)
}

func main() {
	urls := []string{
		"https://github.com",
		"https://github.com/about",
		"https://github.com/contact",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg)
	}

	wg.Wait()
	fmt.Println("All requests completed.")
}
