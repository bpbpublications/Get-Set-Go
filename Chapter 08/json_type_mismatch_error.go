jsonData := `{"name":"Alice","age":"thirty"}`
var p Person

err := json.Unmarshal([]byte(jsonData), &p)
if err != nil {
    fmt.Println("Error decoding JSON:", err)
    fmt.Println("Error type:", reflect.TypeOf(err))
}
