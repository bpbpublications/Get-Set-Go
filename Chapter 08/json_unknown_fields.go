jsonData := `{"name":"Alice","age":30,"gender":"female"}`
var p Person

dec := json.NewDecoder(strings.NewReader(jsonData))
dec.DisallowUnknownFields()

err := dec.Decode(&p)
if err != nil {
    fmt.Println("Error decoding JSON:", err)
    fmt.Println("Error type:", reflect.TypeOf(err))
}
