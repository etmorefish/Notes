package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	r.Run(":8080")
}

/*
curl --location --request POST '127.0.0.1:8080/post?id=123&page=2' \
--form 'name="xxml"' \
--form 'message="helloworld"'
*/
