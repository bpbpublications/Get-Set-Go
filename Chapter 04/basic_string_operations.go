package main

import (
 "fmt"
 "strings"
)

func main() {
 // Checking if a string contains a substring
 fmt.Println("Contains example:")
 fmt.Println(strings.Contains("Hello, Go!", "Go")) // Output: true

 // Counting occurrences of a substring
 fmt.Println("\nCount example:")
 fmt.Println(strings.Count("Hello, Go!", "o")) // Output: 2

 // Replacing substrings
 fmt.Println("\nReplaceAll example:")
 fmt.Println(strings.ReplaceAll("Hello, Go!", "Go", "Golang")) // Output: Hello, Golang!

 // Splitting a string
 fmt.Println("\nSplit example:")
 words := strings.Split("Hello, Go!", ", ")
 fmt.Println(words) // Output: [Hello Go!]

 // Joining a slice of strings
 fmt.Println("\nJoin example:")
 fmt.Println(strings.Join(words, ", ")) // Output: Hello, Go!
}
