package main

import (
 "fmt"
 "reflect"
)

func main() {
 numbers := []int{1, 2, 3, 4}
 modifySlice(&numbers, 2, 99)
 fmt.Println("Modified slice:", numbers)
}

func modifySlice(slice interface{}, index int, newValue interface{}) {
 val := reflect.ValueOf(slice)

 // Ensure it's a pointer to a slice
 if val.Kind() == reflect.Ptr && val.Elem().Kind() == reflect.Slice {
  sliceVal := val.Elem()
  elem := sliceVal.Index(index)

  // Make sure the element is settable and types match
  if elem.CanSet() && reflect.TypeOf(newValue).AssignableTo(elem.Type()) {
   elem.Set(reflect.ValueOf(newValue))
  } else {
   fmt.Println("Cannot set element at index", index)
  }
 } else {
  fmt.Println("Expected a pointer to a slice")
 }
}
