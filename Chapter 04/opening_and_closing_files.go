package main

import (
 "fmt"
 "os"
)

func main() {
 // Opening a file
 file, err := os.Open("example.txt")
 if err != nil {
  fmt.Println("Error opening file:", err)
  return
 }
 // Ensure the file is closed after we are done with it
 defer file.Close()

 // Creating a file
 newFile, err := os.Create("newfile.txt")
 if err != nil {
  fmt.Println("Error creating file:", err)
  return
 }
 defer newFile.Close()

 fmt.Println("Files opened and created!")
}
