package main

import (
 "context"
 "fmt"
 "time"
)

func main() {
 // Create a cancellable context
 ctx, cancel := context.WithCancel(context.Background())
 defer cancel()

 // Launch a goroutine with the context
 go func() {
  for {
   select {
   case <-ctx.Done():
    fmt.Println("Goroutine stopped")
    return
   default:
    fmt.Println("Goroutine running...")
    time.Sleep(500 * time.Millisecond)
   }
  }
 }()

 time.Sleep(2 * time.Second) // Allow the goroutine to run for a bit
 cancel()                    // Cancel the context
 time.Sleep(1 * time.Second)  // Wait for the final message
}
