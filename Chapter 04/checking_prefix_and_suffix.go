package main

import (
 "fmt"
 "strings"
)

func main() {
 fmt.Println(strings.HasPrefix("golang.org", "go")) // Output: true
 fmt.Println(strings.HasSuffix("golang.org", ".org")) // Output: true
}
