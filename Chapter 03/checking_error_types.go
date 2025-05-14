package main

import (
    "errors"
    "fmt"
)

type DivideError struct {
    Dividend float64
    Divisor  float64
}

func (e *DivideError) Error() string {
    return fmt.Sprintf("cannot divide %.2f by %.2f", e.Dividend, e.Divisor)
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, &DivideError{Dividend: a, Divisor: b}
    }
    return a / b, nil
}

func main() {
    _, err := divide(4, 0)
    if err != nil {
        var divideErr *DivideError
        if errors.As(err, &divideErr) {
            fmt.Println("Caught a DivideError:", divideErr)
        }
    }
}
