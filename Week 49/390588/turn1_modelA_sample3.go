package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Format("2006-01-02T15:04:05Z07:00"))
}

func (t *Time) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	parsedTime, err := time.Parse("2006-01-02T15:04:05Z07:00", s)
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}

func main() {
	t := Time{time.Now()}
	b, _ := t.MarshalJSON()
	fmt.Println(string(b))
	t.UnmarshalJSON(b)
}
