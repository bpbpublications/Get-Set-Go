type Employee struct {
    Name     string
    Age      int
    Skills   []string  // Slice of strings
    IsActive bool      // Boolean field
    Promote  func()    // Function field
}
