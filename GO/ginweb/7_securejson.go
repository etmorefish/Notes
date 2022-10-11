package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"foo", "bar", "baz"}
		c.SecureJSON(http.StatusOK, names) // while(1);["foo","bar","baz"]%
	})
	r.Run(":8080")
}
