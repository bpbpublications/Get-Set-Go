package main

import (
    "encoding/json"
    "net/http"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func personHandler(w http.ResponseWriter, r *http.Request) {
    person := Person{Name: "Alice", Age: 30}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(person)
}

func main() {
    http.HandleFunc("/person", personHandler)
    http.ListenAndServe(":8080", nil)
}
