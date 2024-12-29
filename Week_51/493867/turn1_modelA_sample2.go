package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type History struct {
	ModificationTime time.Time
	Author           string
	Changes          []string
}

func logHistory(file string, history History) error {
	historyFileName := fmt.Sprintf("%s.history.json", file)
	data, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(historyFileName, data, 0644)
	return err
}
