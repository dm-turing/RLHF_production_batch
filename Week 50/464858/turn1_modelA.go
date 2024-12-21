package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func evenFilter(src io.Reader, dst io.Writer) error {
	scanner := bufio.NewScanner(src)
	writer := bufio.NewWriter(dst)
	defer writer.Flush()

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Printf("Error converting to int: %v", err)
			continue
		}
		if num%2 == 0 {
			_, err := writer.WriteString(fmt.Sprintf("%d\n", num))
			if err != nil {
				return err
			}
		}
	}

	return scanner.Err()
}

func main() {
	// Create a pipe
	pr, pw := io.Pipe()

	// Start a goroutine to read from the pipe and filter even numbers
	go func() {
		err := evenFilter(pr, os.Stdout)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Write numbers to the pipe
	writer := bufio.NewWriter(pw)
	_, err := writer.WriteString("1 2 3 4 5 6\n")
	if err != nil {
		log.Fatal(err)
	}
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}

	// Close the pipe writer to signal the reader to finish
	err = pw.Close()
	if err != nil {
		log.Fatal(err)
	}
}
