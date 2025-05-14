package main

import "fmt"

func main() {
 str := "नमस्ते, Go!"
 for i, r := range str {
  fmt.Printf("Character %d: %c (Unicode: %U)\n", i, r, r)
 }
}
