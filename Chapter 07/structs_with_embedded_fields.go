type Student struct {
    Person
    StudentID string
}

student := Student{
    Person:    Person{Name: "Bob", Age: 20},
    StudentID: "S1234",
}
fmt.Println(student.Name)     // Output: Bob
fmt.Println(student.StudentID) // Output: S1234
