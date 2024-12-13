package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/errors"
)

func withFile(ctx context.Context, filename string, f func(context.Context, *os.File) ([]byte, error)) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("error closing file: %v", err)
		}
	}()

	return f(ctx, file)
}

func readFileWithContext(ctx context.Context, filename string) ([]byte, error) {
	return withFile(ctx, filename, func(ctx context.Context, file *os.File) ([]byte, error) {
		data, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read file")
		}
		return data, nil
	})
}

func main() {
	ctx := context.Background()
	data, err := readFileWithContext(ctx, "non_existent_file.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
