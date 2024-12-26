package main

import (
    "fmt"
    "strings"
)

// Function signature for a predicate that takes a string and returns a bool
type Predicate func(string) bool

// Function signature for a callback function that takes no parameters and returns no values
type Callback func()

// Function to call the specified callbacks based on the predicate
func callByPredicate(callbacks []Callback, predicate Predicate) {
    for _, callback := range callbacks {
        // Apply the predicate function to the name of the callback function (as a string)
        name := fmt.Sprintf("%v", callback)
        if predicate(name) {
            callback() // Call the callback function
        }
    }
}

// Named function to say "Hello"
func sayHello() {
    fmt.Println("Hello!")
}

// Anonymous function to say "Hi"
hiFn := func() {
    fmt.Println("Hi!")
}

// Named function to say "Greetings"
func sayGreetings() {
    fmt.Println("Greetings!")
}

func main() {
    // Create a slice of callback functions
    callbacks := []Callback{
        sayHello,        // Named function
        hiFn,           // Anonymous function
        sayGreetings,     // Named function
    }

    // Predicate function to check if the function name contains "Hello"
    predicateContainsHello := func(name string) bool {
        return strings.Contains(name, "Hello")
    }

    fmt.Println("Calling functions that contain 'Hello':")
    callByPredicate(callbacks, predicateContainsHello) // Output: Hello! Greetings!

    // Predicate function to check if the function name is even length
    predicateEvenLength := func(name string) bool {
        return len(name)%2 == 0
    }

    fmt.Println("\nCalling functions with even-length names:")
    callByPredicate(callbacks, predicateEvenLength) // Output: Hi! Greetings!
} 