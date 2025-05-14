func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result) // Output: Result: 5
    }

    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)     // Output: Error: cannot divide by zero
    } else {
        fmt.Println("Result:", result)
    }
}
