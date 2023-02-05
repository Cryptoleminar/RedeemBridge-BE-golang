package main

import (
	"github.com/gin-gonic/gin"
)

func example() {
	r := gin.Default()

	// GET request
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// POST request
	r.POST("/post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	// PUT request
	r.PUT("/put", func(c *gin.Context) {
		id := c.Query("id")
		message := c.PostForm("message")

		c.JSON(200, gin.H{
			"id":      id,
			"message": message,
		})
	})

	// DELETE request
	r.DELETE("/delete", func(c *gin.Context) {
		id := c.Query("id")

		c.JSON(200, gin.H{
			"id": id,
		})
	})

	// start server
	r.Run()
}
