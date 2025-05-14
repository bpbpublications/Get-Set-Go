package main

import (
    "context"
    "fmt"
    "net/http"
    "time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    select {
    case <-time.After(2 * time.Second):
        fmt.Fprintln(w, "Hello, World!")
    case <-ctx.Done():
        fmt.Fprintln(w, "Request cancelled")
    }
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.ListenAndServe(":8080", nil)
}
