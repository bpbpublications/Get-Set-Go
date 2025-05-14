var capitals map[string]string = map[string]string{
    "France": "Paris",
    "Japan":  "Tokyo",
}

capitals["Germany"] = "Berlin"
fmt.Println(capitals["France"])  // Output: Paris
fmt.Println(capitals["Germany"]) // Output: Berlin

city, exists := capitals["Italy"]
if exists {
    fmt.Println(city)
} else {
    fmt.Println("Key not found")
}
