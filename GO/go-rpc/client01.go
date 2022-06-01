package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 传的参数
type Params struct {
	Width, Height int
}

// 主函数
func main() {
	// 1.连接远程rpc服务
	conn, err := rpc.DialHTTP("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}
	// 2.调用方法
	// 面积
	ret := 0
	err2 := conn.Call("Rect.Area", Params{50, 100}, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("面积：", ret)
	// 周长
	err3 := conn.Call("Rect.Perimeter", Params{50, 100}, &ret)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println("周长：", ret)
}
