employee := Employee{
    Name:     "Bob",
    Age:      28,
    Skills:   []string{"Go", "Docker", "Kubernetes"},
    IsActive: true,
    Promote: func() {
        fmt.Println("Employee promoted!")
    },
}

fmt.Println(employee.Name)        // Output: Bob
fmt.Println(employee.Skills[0])   // Output: Go
fmt.Println(employee.IsActive)    // Output: true
employee.Promote()    // Output: Employee promoted!
