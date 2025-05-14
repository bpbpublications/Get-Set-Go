func (p Person) DisplayAge() {
    fmt.Printf("I am %d years old.\n", p.Age)
}

person := Person{Name: "Alice", Age: 30}
person.DisplayAge() // Output: I am 30 years old.
