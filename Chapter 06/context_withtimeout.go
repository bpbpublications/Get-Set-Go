package main

import (
 "context"
 "fmt"
 "time"
)

func main() {
 // Create a context with a 2-second timeout
 ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
 defer cancel()

 // Launch a task that will check for context timeout
 go func() {
  select {
  case <-ctx.Done():
   fmt.Println("Timeout reached:", ctx.Err())
   return
  case <-time.After(3 * time.Second): // Simulate long-running task
   fmt.Println("Task completed")
  }
 }()

 time.Sleep(3 * time.Second) // Wait to observe the output
}
