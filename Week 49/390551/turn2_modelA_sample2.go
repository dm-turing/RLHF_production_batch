package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func createFile(filename string) error {
	// Set up zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(os.Stderr)

	file, err := os.Create(filename)
	if err != nil {
		log.Error().Str("filename", filename).Err(err).Msg("Failed to create file")
		return fmt.Errorf("createFile error: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Error().Str("filename", filename).Err(err).Msg("Failed to close file")
		}
	}()

	if _, err := file.WriteString("Hello, world!"); err != nil {
		log.Error().Str("filename", filename).Err(err).Msg("Failed to write to file")
		return fmt.Errorf("writeFile error: %w", err)
	}

	return nil
}

func main() {
	err := createFile("example.txt")
	if err != nil {
		fmt.Println("Error happened:", err)
	}
}
