package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("output.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    fmt.Fprintf(file, "This text is written to a file!\n")
}
