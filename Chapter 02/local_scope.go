func func1() {
    var localVar int = 20
    fmt.Println(localVar) // Accessible
}

func func2() {
    fmt.Println(localVar) // Error: localVar is not accessible here
}
