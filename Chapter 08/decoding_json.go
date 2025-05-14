package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    // JSON string
    jsonData := `{"name":"Alice","age":30}`

    // Define a variable to hold the decoded data
    var p Person

    // Decode JSON into the Go struct
    err := json.Unmarshal([]byte(jsonData), &p)
    if err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }

    // Print the decoded struct
    fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}
