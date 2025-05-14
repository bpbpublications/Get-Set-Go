package main

import (
    "log"
    "os"
)

func main() {
    data, err := os.ReadFile("example.txt")
    if err != nil {
        log.Fatal(err)
    }

    log.Println(string(data))

    msg := []byte("\nWriting data to the file")

    data = append(data, msg...)

    err = os.WriteFile("new_file.txt", data, 0744)
    if err != nil {
        log.Fatal(err)
    }
}
