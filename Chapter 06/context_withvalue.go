package main

import (
 "context"
 "fmt"
)

func main() {
 // Create a context with a key-value pair
 ctx := context.WithValue(context.Background(), "userID", 42)

 // Start a goroutine that accesses the context value
 go func(ctx context.Context) {
  userID := ctx.Value("userID")
  fmt.Println("UserID from context:", userID)
 }(ctx)
}
