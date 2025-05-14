package main

import (
 "fmt"
 "time"
)

func main() {
 // Create an unbuffered channel
 messageChannel := make(chan string)

 // Sender goroutine
 go func() {
  fmt.Println("Sending message...")
  messageChannel <- "Hello, World!" // This will block until the receiver is ready
  fmt.Println("Message sent.")
 }()

 // Simulate a delay before receiving
 time.Sleep(2 * time.Second)

 // Receiver in the main goroutine
 fmt.Println("Receiving message...")
 message := <-messageChannel // This will unblock the sender
 fmt.Printf("Message received: %s\n", message)
}
