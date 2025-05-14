package main

import (
 "fmt"
 "net/http"

 "github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func loginHandler(w http.ResponseWriter, r *http.Request) {
 // Simulate login (replace with actual authentication logic)
 username := r.FormValue("username")
 if username == "admin" {
  session, _ := store.Get(r, "session-id")
  // Set session values
  session.Values["username"] = username
  session.Save(r, w)
  fmt.Fprintln(w, "Login successful!")
  return
 }
 fmt.Fprintln(w, "Invalid username")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
 session, _ := store.Get(r, "session-id")
 // Remove session values
 session.Options.MaxAge = -1 // Mark session as expired
 session.Save(r, w)
 fmt.Fprintln(w, "Logged out successfully!")
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
 session, _ := store.Get(r, "session-id")
 // Check if user is logged in
 if username, ok := session.Values["username"].(string); ok {
  fmt.Fprintf(w, "Welcome, %s!", username)
  return
 }
 http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func main() {
 http.HandleFunc("/login", loginHandler)
 http.HandleFunc("/logout", logoutHandler)
 http.HandleFunc("/protected", protectedHandler)

 fmt.Println("Server running on http://localhost:8080")
 http.ListenAndServe(":8080", nil)
}
