package main

import (
 "fmt"
 "net/http"
 "path"
 "html/template"
)

// indexHandler serves the HTML form for GET requests.
func indexHandler(w http.ResponseWriter, r *http.Request) {
 fp := path.Join("templates", "index.html")
 tmpl, err := template.ParseFiles(fp)
 if err != nil {
  http.Error(w, err.Error(), http.StatusInternalServerError)
  return
 }

 if err := tmpl.Execute(w, nil); err != nil {
  http.Error(w, err.Error(), http.StatusInternalServerError)
 }
}

// formHandler processes POST requests for the form.
func formHandler(w http.ResponseWriter, r *http.Request) {
 if r.Method == http.MethodPost {
  // Retrieve the value of the "name" field from the form
  name := r.FormValue("name")
  if name == "" {
   // Handle empty input with a message
   http.Error(w, "Name field is required", http.StatusBadRequest)
   return
  }
  // Respond with a personalized greeting
  fmt.Fprintf(w, "Hello, %s!", name)
  return
 }
 http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}

func main() {
 http.HandleFunc("/", indexHandler)

 // Map the /form route to the formHandler
 http.HandleFunc("/form", formHandler)

 // Start the HTTP server on port 8080
 fmt.Println("Server running on http://localhost:8080/form")
 if err := http.ListenAndServe(":8080", nil); err != nil {
  fmt.Println("Error starting server:", err)
 }
}
