var data map[string]interface{}
err := json.Unmarshal([]byte(jsonData), &data)
if err != nil {
    fmt.Println("Error decoding JSON:", err)
    return
}
fmt.Println(data["name"])
