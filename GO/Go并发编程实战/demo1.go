package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(map[int]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	for k, v := range data {
		//下面会出现单个值打印多次，由于变量名没有刷新
		// go func() {
		// 	fmt.Println("k->", k, "v->", v)
		// }()

		// 传值强制刷新
		go func(k int, v int) {
			fmt.Println("k->", k, "v->", v)
		}(k, v)
	}
	time.Sleep(time.Second * 2)
}
