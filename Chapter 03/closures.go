func main() {
    x := 5
    increment := func() int {
        x = x + 1
        return x
    }
    fmt.Println(increment()) // Output: 6
    fmt.Println(increment()) // Output: 7
}
