package main

import (
 "fmt"
 "net/http"
)

func methodHandler(w http.ResponseWriter, r *http.Request) {
 switch r.Method {
 case http.MethodGet:
  fmt.Fprintln(w, "This is a GET request")
 case http.MethodPost:
  fmt.Fprintln(w, "This is a POST request")
 default:
  http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
 }
}

func main() {
 http.HandleFunc("/method", methodHandler)
 http.ListenAndServe(":8080", nil)
}
