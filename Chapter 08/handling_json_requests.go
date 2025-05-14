package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func personHandler(w http.ResponseWriter, r *http.Request) {
    var person Person
    err := json.NewDecoder(r.Body).Decode(&person)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    fmt.Fprintf(w, "Received: %s, Age: %d\n", person.Name, person.Age)
}

func main() {
    http.HandleFunc("/person", personHandler)
    http.ListenAndServe(":8080", nil)
}
