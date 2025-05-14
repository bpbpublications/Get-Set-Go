func subtract(a, b int) int {
    return a - b
}

func divide(a, b float64) float64 {
    if b == 0 {
        return 0
    }
    return a / b
}

func greet(name string, age int) {
    fmt.Printf("Hello %s, you are %d years old.\n", name, age)
}