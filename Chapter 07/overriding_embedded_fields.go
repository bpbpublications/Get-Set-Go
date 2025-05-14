type Graduate struct {
    Person
    Degree string
    Name   string // Overrides Person.Name
}

graduate := Graduate{
    Person: Person{Name: "Alice", Age: 25},
    Degree: "Master's",
    Name:   "Dr. Alice",
}
fmt.Println(graduate.Name)  // Output: Dr. Alice
fmt.Println(graduate.Person.Name) // Output: Alice
