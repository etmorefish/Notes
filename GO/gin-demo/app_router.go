package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	r.POST("carts", api.CreateCart)    // POST请求，一般提交表单
	r.GET("carts/:id", api.ShowCarts)  //GET请求，一般获取数据
	r.PUT("carts", api.UpdateCart)     //PUT请求，一般修改数据
	r.DELETE("carts", api.DeleteCart)  // DELETE请求，一般删除数据

*/

func main() {
	r := gin.Default() // default to create router
	r.GET("/ping", func(c *gin.Context) {
		// c.String(http.StatusOK, "OK") // return status code and body
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// GET DELETE accept params => .Param
	r.GET("/user/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.JSON(200, gin.H{
			"name":   name,
			"action": action,
		})
	})

	// POST PUT
	// Form data  => .PostForm
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(200, gin.H{
			"username": username,
			"password": password,
		})
	})

	// upload file => .FormFile
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(500, gin.H{
				"status_code": http.StatusBadRequest,
				"error":       err,
			})
		}
		c.SaveUploadedFile(file, "./"+file.Filename)
		c.JSON(200, gin.H{
			"status_code": http.StatusOK,
			"filename":    file.Filename,
		})
	})

	// // download
	// // In general, the route will be split into URLs and processing functions,
	// //  which can effectively manage the logic of our framework.
	// func helloHandler(c *gin.Context){
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Hello World!",
	// 	})
	// }

	// r.GET("/hello", helloHandler)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
