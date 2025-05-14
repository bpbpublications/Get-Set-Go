package main

import (
 "fmt"
 "net/http"
 "strings"
 "time"

 "github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your-secret-key")

func generateJWT(username string) (string, error) {
	// Create a new token object with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateJWT(tokenString string) (string, error) {
 // Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is as expected
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	// Extract and verify claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	}

	return "", fmt.Errorf("invalid token")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username != "exampleUser" || password != "password123" {
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := generateJWT(username)
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"token": "%s"}`, token)
}

func secureHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "missing token", http.StatusUnauthorized)
		return
	}

	// Extract token from the "Bearer <token>" format
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	username, err := validateJWT(tokenString)
	if err != nil {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Welcome, %s!", username)
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/secure", secureHandler)
	http.ListenAndServe(":8080", nil)
}
