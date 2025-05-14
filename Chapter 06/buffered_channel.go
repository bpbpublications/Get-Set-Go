package main

import (
 "fmt"
 "time"
)

func main() {
 // Create a buffered channel with a capacity of 3
 messageChannel := make(chan string, 3)

 // Sender goroutine
 go func() {
  fmt.Println("Sending message 1")
  messageChannel <- "Message 1" // Won't block
  fmt.Println("Sending message 2")
  messageChannel <- "Message 2" // Won't block
  fmt.Println("Sending message 3")
  messageChannel <- "Message 3" // Won't block
  fmt.Println("All messages sent.")
 }()

 // Simulate a delay before receiving
 time.Sleep(2 * time.Second)

 // Receiver in the main goroutine
 for i := 1; i <= 3; i++ {
  fmt.Printf("Receiving message %d\n", i)
  fmt.Printf("Message received: %s\n", <-messageChannel)
 }
}
