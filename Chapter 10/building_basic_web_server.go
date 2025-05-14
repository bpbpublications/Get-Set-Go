package main

import (
 "fmt"
 "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintf(w, "Welcome to the Go Web Server!")
}

func main() {
 http.HandleFunc("/", homePage)
 fmt.Println("Server starting on port 8080...")
 http.ListenAndServe(":8080", nil)
}
