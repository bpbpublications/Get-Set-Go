package main

import "fmt"

// Container defines an interface for types that can hold elements of a specific type.
type Container[T any] interface {
    Add(T)
    Get() (T, bool)
}

// MySlice implements the Container interface.
type MySlice[T any] []T

func (ms *MySlice[T]) Add(val T) {
    *ms = append(*ms, val)
}

func (ms *MySlice[T]) Get() (T, bool) {
    if len(*ms) == 0 {
        var zero T
        return zero, false
    }
    val := (*ms)[0]
    *ms = (*ms)[1:]
    return val, true
}

func main() {
    var intContainer Container[int] = &MySlice[int]{}
    intContainer.Add(10)
    val, ok := intContainer.Get()
    if ok {
        fmt.Println(val) // Output: 10
    }

    var stringContainer Container[string] = &MySlice[string]{}
    stringContainer.Add("Hello")
    strVal, ok := stringContainer.Get()
    if ok {
        fmt.Println(strVal) // Output: Hello
    }
}
