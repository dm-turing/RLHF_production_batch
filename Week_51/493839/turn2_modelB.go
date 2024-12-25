package main

import (
	"fmt"
	"strings"
	"time"
)

func logMessage(level string, message string, data map[string]interface{}) {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[%s] %s: ", time.Now().Format(time.RFC3339), level))
	sb.WriteString(message)
	sb.WriteString(" ")
	for key, value := range data {
		sb.WriteString(fmt.Sprintf("%s=%v ", key, value))
	}
	sb.WriteByte('\n')
	fmt.Print(sb.String())
}

func main() {
	logMessage("INFO", "Application started", map[string]interface{}{"port": 8080, "env": "production"})
}
