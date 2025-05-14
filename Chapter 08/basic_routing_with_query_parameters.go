package main

import (
 "fmt"
 "net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
 name := r.URL.Query().Get("name")
 if name == "" {
  name = "World"
 }
 fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
 http.HandleFunc("/greet", greetHandler)
 http.ListenAndServe(":8080", nil)
}
