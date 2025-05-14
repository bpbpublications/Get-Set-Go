package main

import (
 "fmt"
 "sync"
)

func main() {
 var once sync.Once
 var wg sync.WaitGroup

 initialize := func() {
  fmt.Println("Initializing resources")
 }

 for i := 1; i <= 3; i++ {
  wg.Add(1)
  go func(id int) {
   defer wg.Done()
   once.Do(initialize) // Ensures initialize runs only once
   fmt.Printf("Goroutine %d completed\n", id)
  }(i)
 }

 wg.Wait()
}
