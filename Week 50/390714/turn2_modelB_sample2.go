package main

import (
	"errors"
	"fmt"
)

// Fallback strategies
const (
	FallbackToDatabase = iota
	FallbackToCache    = iota
	FallbackToDefault  = iota
)

// Sample database and cache services
var (
	ErrDatabaseDown = errors.New("database is down")
	ErrCacheMiss    = errors.New("cache miss")
)

// Simulating database and cache operations
func databaseOperation() error {
	//Fake down scenario
	return ErrDatabaseDown
}

func cacheOperation() error {
	//Fake miss scenario
	return ErrCacheMiss
}

// Decision tree to select fallback strategy
func selectFallback() int {
	if databaseOperation() == nil {
		return FallbackToDatabase
	}
	if cacheOperation() == nil {
		return FallbackToCache
	}
	return FallbackToDefault
}

// Main function to demonstrate the usage
func main() {
	fallbackStrategy := selectFallback()
	switch fallbackStrategy {
	case FallbackToDatabase:
		fmt.Println("Using Database as fallback.")
	case FallbackToCache:
		fmt.Println("Using Cache as fallback.")
	case FallbackToDefault:
		fmt.Println("Using Default fallback.")
	}
}
