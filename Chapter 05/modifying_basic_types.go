package main

import (
 "fmt"
 "reflect"
)

func main() {
 var x int = 10
 var y string = "Hello"

 modifyValue(&x, 20)
 modifyValue(&y, "World")

 fmt.Println("Modified int:", x)    // Modified int: 20
 fmt.Println("Modified string:", y) // Modified string: World
}

func modifyValue(variable interface{}, newValue interface{}) {
 val := reflect.ValueOf(variable)

 // Check if the reflect.Value is a pointer and is settable
 if val.Kind() == reflect.Ptr && val.Elem().CanSet() {
  newVal := reflect.ValueOf(newValue)

  // Make sure the types match before setting the value
  if newVal.Type().AssignableTo(val.Elem().Type()) {
   val.Elem().Set(newVal)
  } else {
   fmt.Println("Type mismatch")
  }
 } else {
  fmt.Println("Value is not settable")
 }
}
