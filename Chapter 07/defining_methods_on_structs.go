type Person struct {
    Name string
    Age  int
}

// Method with a value receiver
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s.\n", p.Name)
}
