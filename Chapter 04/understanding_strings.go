package main

import "fmt"

func main() {
 str := "Hello"
 // Attempting to modify string directly causes an error
 // str[0] = 'h' // This throws a compile-time error

 // Instead, we create a new string
 newStr := "h" + str[1:]
 fmt.Println(newStr) // Output: "hello"
}
