package main

import (
 "fmt"
 "reflect"
)

func printTypeAndValue(i interface{}) {
 t := reflect.TypeOf(i)
 v := reflect.ValueOf(i)
 fmt.Printf("Type: %s, Value: %v\n", t, v)
}

func main() {
 printTypeAndValue(42)
 printTypeAndValue("Hello, Go!")
 printTypeAndValue(3.14)
}
