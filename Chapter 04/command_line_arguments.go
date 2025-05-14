package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args // Capture all command-line arguments
    fmt.Println("Program Name:", args[0]) // Print the program name

    if len(args) > 1 {
        fmt.Println("Arguments passed:")
        for i, arg := range args[1:] {
            fmt.Printf("Argument %d: %s\n", i+1, arg)
        }
    } else {
        fmt.Println("No arguments provided.")
    }
}
