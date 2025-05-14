package main

import (
 "net/http"
)

func basicAuth(next http.HandlerFunc) http.HandlerFunc {
 return func(w http.ResponseWriter, r *http.Request) {
  username, password, ok := r.BasicAuth()
  if !ok || username != "admin" || password != "password" {
   w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
   http.Error(w, "Unauthorized", http.StatusUnauthorized)
   return
  }
  next(w, r)
 }
}

func secretPage(w http.ResponseWriter, r *http.Request) {
 w.Write([]byte("Welcome to the secret page!"))
}

func main() {
 http.HandleFunc("/secret", basicAuth(secretPage))
 http.ListenAndServe(":8080", nil)
}
