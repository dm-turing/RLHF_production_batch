package main

import (
	"fmt"
	"strings"
)

func isValidUserId(userId string) bool {
	if len(userId) != 10 || !strings.ContainsFunc(userId, func(r rune) bool {
		return '0' <= r && r <= '9'
	}) {
		return false
	}
	return true
}

func main() {
	userId := "1234567890"
	if !isValidUserId(userId) {
		fmt.Println("Invalid user ID.")
		return
	}
	fmt.Println("Valid user ID.")
}
