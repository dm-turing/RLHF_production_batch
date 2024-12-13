package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"

	"github.com/rs/zerolog"
)

func openFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to open file %s", filename))
	}
	defer func() {
		if err := file.Close(); err != nil {
			zerolog.Ctx(log.Ctx()).Error().Err(err).Msg("Error closing file")
		}
	}()
	return nil
}

func main() {

	zerolog.New(os.Stdout).Level(zerolog.DebugLevel).EnrichWith(zerolog.Caller()).Log().Msg("Application started")

	if err := openFile("non_existent_file.txt"); err != nil {
		zerolog.Ctx(log.Ctx()).Error().Err(err).Msg("Error opening file")
	}

	zerolog.Ctx(log.Ctx()).Msg("Application terminated")
}
