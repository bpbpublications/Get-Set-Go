type Employee struct {
    name string  // unexported field
}

func NewEmployee(name string) Employee {
    return Employee{name: name}
}

func (e Employee) GetName() string {
    return e.name
}
