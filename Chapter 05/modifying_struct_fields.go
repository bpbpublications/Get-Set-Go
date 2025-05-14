package main

import (
 "fmt"
 "reflect"
)

type Person struct {
 Name string
 Age  int
}

func main() {
 p := &Person{Name: "Alice", Age: 30}
 modifyField(p, "Name", "Bob")
 modifyField(p, "Age", 40)

 fmt.Printf("Modified Person: %+v\n", p) // Modified Person: &{Name:Bob Age:40}
}

func modifyField(s interface{}, fieldName string, newValue interface{}) {
 val := reflect.ValueOf(s)

 // Ensure the reflect.Value is a pointer to a struct
 if val.Kind() == reflect.Ptr && val.Elem().Kind() == reflect.Struct {
  field := val.Elem().FieldByName(fieldName)

  // Check if the field is valid and settable
  if field.IsValid() && field.CanSet() {
   newVal := reflect.ValueOf(newValue)

   // Make sure the types are compatible
   if newVal.Type().AssignableTo(field.Type()) {
    field.Set(newVal)
   } else {
    fmt.Printf("Type mismatch: cannot assign %v to field %s\n", newVal.Type(), fieldName)
   }
  } else {
   fmt.Printf("Field %s is not settable or does not exist\n", fieldName)
  }
 } else {
  fmt.Println("Expected a pointer to a struct")
 }
}
