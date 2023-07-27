package main

import "github.com/gin-gonic/gin"

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Define the route
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// Run the server on port 8080
	r.Run(":8080")
}
