package main

import "encoding/json"

func main() {
	var data interface{}
	jsonData := `[{"name":"Alice","age":30},{"name":"Bob","age":25}]`
	json.Unmarshal([]byte(jsonData), &data)
}
