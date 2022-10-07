package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
		fmt.Println()
		for k, v := range ids {
			fmt.Println(k, v)
		}
		for k, v := range names {
			fmt.Println(k, v)
		}
	})
	router.Run(":8080")
}

/*
curl --location -g --request POST '127.0.0.1:8080/post?ids[a]=1234&ids[b]=hello' \
--header 'names;' \
--form 'names[first]="wqwq"' \
--form 'names[second]="adddd"'
*/
