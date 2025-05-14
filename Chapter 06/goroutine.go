package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello, Go!")
}

func main() {
    go sayHello()
    time.Sleep(1 * time.Second)
}
