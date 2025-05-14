func (p *Person) IncrementAge() {
    p.Age++
}

person := Person{Name: "Alice", Age: 30}
person.IncrementAge()
fmt.Println(person.Age) // Output: 31
