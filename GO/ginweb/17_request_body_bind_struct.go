package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func main() {
	r := gin.Default()

	// 简单的路由组: v1
	v1 := r.Group("/v1")
	{
		v1.GET("/some", SomeHandler)
		v1.GET("/some2", SomeHandler2)

	}

	r.Run(":8080")
}

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func SomeHandler(c *gin.Context) {
	objA := formA{Foo: "foooo"}
	objB := formB{Bar: "barrr"}
	// c.ShouldBind 使用了 c.Request.Body，不可重用。
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// 因为现在 c.Request.Body 是 EOF，所以这里会报错。
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		c.String(http.StatusOK, `else`)
	}

}

func SomeHandler2(c *gin.Context) {
	// 要想多次绑定，可以使用 c.ShouldBindBodyWith.

	objA := formA{Foo: "foooo"}
	objB := formB{Bar: "barrr"}
	// 读取 c.Request.Body 并将结果存入上下文。
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// 这时, 复用存储在上下文中的 body。
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, `the body should be formB JSON`)
		// 可以接受其他格式
	} else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
		c.String(http.StatusOK, `the body should be formB XML`)
	} else {
		c.String(http.StatusOK, `else`)

	}
}
