package main

import (
 "fmt"
 "sync"
)

func main() {
 var count int
 var mu sync.Mutex
 var wg sync.WaitGroup

 for i := 0; i < 5; i++ {
  wg.Add(1)
  go func() {
   defer wg.Done()
   mu.Lock() // Lock before accessing resource
   fmt.Printf("Locking Mutex - Goroutine %d\n", i)
   count++ // Safe increment
   fmt.Printf("Unlocking Mutex - Goroutine %d\n", i)
   mu.Unlock() // Unlock after accessing resource
  }()
 }

 wg.Wait()
 fmt.Println("Final Count:", count)
}
