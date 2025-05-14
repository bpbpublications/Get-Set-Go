jsonArray := `[{"name":"Alice","age":30},{"name":"Bob","age":25}]`
var people []Person
err := json.Unmarshal([]byte(jsonArray), &people)
if err != nil {
    fmt.Println("Error decoding JSON array:", err)
    return
}
fmt.Println(people)
