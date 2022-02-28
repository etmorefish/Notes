package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) { // 把处理函数放在这里，可以处理一些复杂的业务
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!", // 通过这个 gin.H{} 进行json格式的返回
	})
}

// Router 配置路由信息
func Router() *gin.Engine {

	r.GET("/hello", HelloHandler) // 这样这个路由就比较简洁了。
	return r
}
func main() {
	r := gin.Default()
	r.run()
}
