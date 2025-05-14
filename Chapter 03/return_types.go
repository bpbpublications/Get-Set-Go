func add(a int, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("Cannot add zero")
    }
    return a + b, nil
}

func splitName(fullName string) (string, string) {
    parts := strings.Split(fullName, " ")
    return parts[0], parts[1]
}

func getUserDetails(id int) (name string, age int, error) {
    // Assume there's logic to retrieve details
    return "John Doe", 30, nil
}
