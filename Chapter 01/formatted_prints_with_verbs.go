package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    fmt.Printf("Integer: %d\n", 100)
    fmt.Printf("Float: %.2f\n", 123.456) // format float with 2 decimal places
    fmt.Printf("String: %s\n", "Go is fun")
    fmt.Printf("Boolean: %t\n", true)

    // Using %+v to display struct with field names
    p := Person{"Alice", 30}
    fmt.Printf("Struct: %+v\n", p)
    fmt.Printf("Type of variable p: %T\n", p)
}
