package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模拟一些私人数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()

	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())

	// 你可以为每个路由添加任意数量的中间件。
	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	// 认证路由组
	// authorized := r.Group("/", AuthRequired())
	// 和使用以下两行代码的效果完全一样:
	authorized = r.Group("/")
	// 路由组中间件! 在此例中，我们在 "authorized" 路由组中使用自定义创建的
	// AuthRequired() 中间件
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		// 嵌套路由组
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

func loginEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "loginEndpoint", "status": http.StatusOK})

}

func readEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "readEndpoint", "status": http.StatusOK})

}

func analyticsEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "analyticsEndpoint", "status": http.StatusOK})

}

func submitEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "submitEndpoint", "status": http.StatusOK})

}

func AuthRequired() gin.HandlerFunc {
	return nil
}

func benchEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "benchEndpoint", "status": http.StatusOK})

}

func MyBenchLogger() gin.HandlerFunc {
	return nil

}
