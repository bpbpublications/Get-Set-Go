package main

import (
  "regexp"
  "testing"
)

func TestValidateEmail(t *testing.T) {
  tests := []struct {
    email    string
    expected bool
  }{
     {"user@example.com", true}, 
     {"invalid-email", false}, 
     {"", false}, 
     {"user@company.org", true}, 
  }

  for _, test := range tests {
    result := ValidateEmail(test.email)
    if result != test.expected {
      t.Errorf("ValidateEmail(%q) = %v; want %v", test.email, result, test.expected)
    }
  }
}

func ValidateEmail(email string) bool {
  // Simple regex pattern for email validation
  emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
  re := regexp.MustCompile(emailRegex)
  return re.MatchString(email)
}
