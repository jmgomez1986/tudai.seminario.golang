package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/users/:name", getUserHandler)
	r.POST("/users", addUserHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getUserHandler(c *gin.Context) {
	name := c.Param("name")
	lastname := c.Query("lastname")
	c.JSON(200, gin.H{
		"status": "ok",
		"name":   name + " " + lastname,
	})
}

func addUserHandler(c *gin.Context) {
	c.JSON(201, gin.H{
		"name":     "matias",
		"lastname": "gomez",
	})
}
