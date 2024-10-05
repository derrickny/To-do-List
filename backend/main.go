package main

import (
	"github.com/gin-gonic/gin"
)

// main is the entry point for the application
func main() {
	// r is the router instance, which is used to route HTTP requests to the appropriate handler functions
	r := gin.Default()

	// This is a GET request to the root URL ("/"), which will return a JSON response
	r.GET("/", func(c *gin.Context) {
		// c is the context object, which contains information about the current request and response
		// gin.H is a helper function that returns a JSON object with the specified key-value pairs
		// In this case, we're returning a JSON object with a single key-value pair: "message" with the value "Welcome to the Go backend For To-Do List"
		c.JSON(200, gin.H{
			"message": "Welcome to the Go backend For To-Do List",
		})
	})

	// This line starts the server and listens on port 8080
	r.Run(":8080")
}
