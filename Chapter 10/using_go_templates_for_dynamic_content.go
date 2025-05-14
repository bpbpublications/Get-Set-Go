package main

import (
 "html/template"
 "log"
 "net/http"
 "path"
)

// homePage is an HTTP handler that serves the homepage.
func homePage(w http.ResponseWriter, r *http.Request) {
 // Dynamic data to render
 data := map[string]string{
  "Title":   "Welcome",
  "Message": "Hello, Go Templates!",
 }

 fp := path.Join("templates", "index.html")
 tmpl, err := template.ParseFiles(fp)
 if err != nil {
  http.Error(w, err.Error(), http.StatusInternalServerError)
  return
 }

 if err := tmpl.Execute(w, data); err != nil {
  http.Error(w, err.Error(), http.StatusInternalServerError)
 }
}

func main() {
 // Map the root URL path to the homePage handler
 http.HandleFunc("/", homePage)

 // Start the HTTP server on port 8080
 log.Println("Server running on localhost:8080")
 err := http.ListenAndServe(":8080", nil)
 if err != nil {
  log.Fatal("ListenAndServe error:", err)
 }
}
