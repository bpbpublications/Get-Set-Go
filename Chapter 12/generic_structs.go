package main

import "fmt"

// Stack is a generic stack data structure.
type Stack[T any] struct {
    elements []T
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(element T) {
    s.elements = append(s.elements, element)
}

// Pop removes and returns the top element from the stack.
// It returns a zero value of T and false if the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
    if s.IsEmpty() {
        var zero T
        return zero, false // Return zero value and false if stack is empty
    }
    element := s.elements[len(s.elements)-1]
    s.elements = s.elements[:len(s.elements)-1]
    return element, true
}

// IsEmpty returns true if the stack is empty, false otherwise.
func (s *Stack[T]) IsEmpty() bool {
    return len(s.elements) == 0
}

func main() {
    intStack := Stack[int]{}
    intStack.Push(1)
    intStack.Push(2)
    intStack.Push(3)

    val, ok := intStack.Pop()
    if ok {
        fmt.Println("Popped int:", val) // Output: Popped int: 3
    }

    val, ok = intStack.Pop()
    if ok {
        fmt.Println("Popped int:", val) // Output: Popped int: 2
    }

    stringStack := Stack[string]{}
    stringStack.Push("hello")
    stringStack.Push("world")

    strVal, ok := stringStack.Pop()
    if ok {
        fmt.Println("Popped string:", strVal) // Output: Popped string: world
    }

    emptyStack := Stack[int]{}
    _, ok = emptyStack.Pop()
    if !ok {
        fmt.Println("Cannot pop from an empty stack") // Output: Cannot pop from an empty stack
    }
}
