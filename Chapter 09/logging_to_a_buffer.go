package main

import (
    "bytes"
    "fmt"
    "log"
)

func main() {
    var buffer bytes.Buffer
    bufferLogger := log.New(&buffer, "[BUFFER] ", log.LstdFlags)
    bufferLogger.Println("This log is written to a buffer")
    fmt.Println(buffer.String())
}
