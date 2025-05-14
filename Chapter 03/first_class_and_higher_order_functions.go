package main

import "fmt"

// First-class function
var greet = func(name string) string {
    return "Hello, " + name
}

// Higher-order function
func processGreeting(f func(string) string, name string) {
    fmt.Println(f(name))
}

func main() {
    // Using the first-class function directly
    message := greet("Spock")
    fmt.Println(message) // Output: Hello, Spock

    // Passing the first-class function to a higher-order function
    processGreeting(greet, "Ripley") // Output: Hello, Ripley
}
