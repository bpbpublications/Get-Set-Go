var i interface{} = 42
n, ok := i.(int)
if ok {
    fmt.Println(n) // Output: 42
} else {
    fmt.Println("Type assertion failed.")
}
