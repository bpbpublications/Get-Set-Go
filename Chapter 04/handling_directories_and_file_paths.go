package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

func main() {
    rootDir := "exampleDir"
    nestedDir := filepath.Join(rootDir, "nestedDir")

    // Step 1: Create a root directory and a nested directory
    if err := os.MkdirAll(nestedDir, 0755); err != nil {
        log.Fatalf("Error creating directories: %v", err)
    }
    fmt.Println("Directories created:", rootDir, "and", nestedDir)

    // Step 2: Create a new file within the nested directory
    filePath := filepath.Join(nestedDir, "file.txt")
    fileContent := []byte("This is a sample file.")
    if err := os.WriteFile(filePath, fileContent, 0644); err != nil {
        log.Fatalf("Error creating file: %v", err)
    }
    fmt.Println("File created:", filePath)

    // Step 3: Get absolute path for the nested directory and the file
    absRootDir, err := filepath.Abs(rootDir)
    if err != nil {
        log.Fatalf("Error getting absolute path: %v", err)
    }
    fmt.Println("Absolute path of root directory:", absRootDir)

    absFilePath, err := filepath.Abs(filePath)
    if err != nil {
        log.Fatalf("Error getting absolute file path: %v", err)
    }
    fmt.Println("Absolute path of file:", absFilePath)

    // Step 4: List all files and directories within the root directory
    fmt.Println("Directory structure:")
    err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        fmt.Println(" -", path)
        return nil
    })
    if err != nil {
        log.Fatalf("Error walking the directory tree: %v", err)
    }

    // Step 5: Clean up - Remove all directories and files created
    if err := os.RemoveAll(rootDir); err != nil {
        log.Fatalf("Error removing directories: %v", err)
    }
    fmt.Println("Directories and files cleaned up.")
}
