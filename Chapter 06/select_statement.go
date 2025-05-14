package main

import (
 "fmt"
 "time"
)

func main() {
 // Create two channels
 ch1 := make(chan string)
 ch2 := make(chan string)

 // Start goroutine to send to ch1
 go func() {
  for {
   ch1 <- "Message from Channel 1"
   time.Sleep(1 * time.Second)
  }
 }()

 // Start goroutine to send to ch2
 go func() {
  for {
   ch2 <- "Message from Channel 2"
   time.Sleep(2 * time.Second)
  }
 }()

 // Listen to both channels using select
 for i := 0; i < 5; i++ { // Limit to 5 messages for example purposes
  select {
  case msg1 := <-ch1:
   fmt.Println("Received:", msg1)
  case msg2 := <-ch2:
   fmt.Println("Received:", msg2)
  }
 }

 fmt.Println("Finished listening to channels")
}
