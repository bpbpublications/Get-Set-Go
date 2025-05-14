package main

import (
    "log"
    "os"
)

func main() {
    customLogger := log.New(os.Stdout, "[CUSTOM] ", log.LstdFlags|log.Lmsgprefix)
    customLogger.Println("Logging with a custom prefix")
}
