package main

import (
	"encoding/json"
	"fmt"
)

type SensitiveData struct {
	Secret string `json:"secret"`
}

func (sd *SensitiveData) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{"secret": "********"})
}

func (sd *SensitiveData) UnmarshalJSON(data []byte) error {
	var temp struct {
		Secret string `json:"secret"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	sd.Secret = temp.Secret
	return nil
}

func main() {
	t := SensitiveData{Secret: "secret_key"}
	b, _ := t.MarshalJSON()
	fmt.Println(string(b))
	t.UnmarshalJSON(b)
	fmt.Println(t.Secret)
}
