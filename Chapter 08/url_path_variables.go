package main

import (
 "fmt"
 "net/http"
 "strings"
)

func dynamicHandler(w http.ResponseWriter, r *http.Request) {
 id := strings.TrimPrefix(r.URL.Path, "/user/")
 if id == "" {
  http.Error(w, "User ID not provided", http.StatusBadRequest)
  return
 }
 fmt.Fprintf(w, "User ID: %s", id)
}

func main() {
 http.HandleFunc("/user/", dynamicHandler)
 http.ListenAndServe(":8080", nil)
}
