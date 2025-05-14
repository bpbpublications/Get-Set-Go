package main

import "fmt"

// Signed is a constraint that permits any signed integer type.
type Signed interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Abs returns the absolute value of x.
func Abs[T Signed](x T) T {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    fmt.Println(Abs(-5))    // Output: 5
    fmt.Println(Abs(int8(-10))) // Output: 10
    fmt.Println(Abs(int64(-15))) // Output: 15
}
