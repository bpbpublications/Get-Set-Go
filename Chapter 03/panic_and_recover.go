func main() {
    // defer ensures this function runs even after a panic
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()

    fmt.Println("Starting the program")
    panic("Something went wrong!") // triggers a panic
    fmt.Println("This will not be printed") // unreachable due to panic
}
