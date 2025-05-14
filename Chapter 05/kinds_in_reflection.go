package main

import (
 "fmt"
 "reflect"
)

func PrintValue(v interface{}) {
 val := reflect.ValueOf(v)
 switch val.Kind() {
 case reflect.Int:
  fmt.Printf("Integer: %d\n", val.Int())
 case reflect.String:
  fmt.Printf("String: %s\n", val.String())
 case reflect.Slice:
  fmt.Printf("Slice: %v\n", val.Interface())
 default:
  fmt.Println("Unsupported type")
 }
}

func main() {
 PrintValue(42)
 PrintValue("Hello, Go!")
 PrintValue([]int{1, 2, 3})
}
