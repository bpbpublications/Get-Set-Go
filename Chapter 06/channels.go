package main

import (
 "fmt"
 "sync"
 "time"
)

func main() {
 var wg sync.WaitGroup // Declare a WaitGroup

 // Create a channel of type int
 messageChannel := make(chan int)

 wg.Add(2) // We have two goroutines to wait for

 // Start a goroutine to send data into the channel
 go func() {
  defer wg.Done() // Mark this goroutine as done when it completes
  for i := 1; i <= 5; i++ {
   fmt.Printf("Sending message %d\n", i)
   messageChannel <- i         // Send the value to the channel
   time.Sleep(1 * time.Second) // Simulate delay
  }
  close(messageChannel) // Close the channel after sending all messages
 }()

 // Start a goroutine to receive data from the channel
 go func() {
  defer wg.Done()                   // Mark this goroutine as done when it completes
  for msg := range messageChannel { // Range over the channel to receive messages
   fmt.Printf("Received message %d\n", msg)
  }
 }()

 // Wait for both goroutines to complete
 wg.Wait()
 fmt.Println("Program finished")
}
