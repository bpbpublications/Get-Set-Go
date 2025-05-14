file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
if err != nil {
    log.Fatal("Error opening log file:", err)
}
defer file.Close()
fileLogger := log.New(file, "[FILE] ", log.LstdFlags)
fileLogger.Println("This log is written to a file")
