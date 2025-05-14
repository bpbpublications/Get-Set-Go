package main

import (
 "github.com/gin-gonic/gin"
)

func main() {
 // Create a new Gin router with default middleware
 router := gin.Default()

 // Define a route for the root URL
 router.GET("/", func(ctx *gin.Context) {
  ctx.String(200, "Welcome to Gin!")
 })

 // Start the server on port 8080
 router.Run(":8080")
}
