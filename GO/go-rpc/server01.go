package main

import (
	"log"
	"net/http"
	"net/rpc"
)

//    例题：golang实现RPC程序，实现求矩形面积和周长
type Params struct {
	Width, Height int //
}

// 结构体，用于注册的
type Rect struct{}

// RPC服务端方法，求矩形面积
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}

// 周长
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Width) * 2
	return nil
}

func main() {
	// 1.注册服务
	rect := new(Rect)
	// 注册一个rect的服务
	rpc.Register(rect)
	// 2.服务处理绑定到http上
	rpc.HandleHTTP()
	// 3.监听服务
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Panicln(err)
	}
}
