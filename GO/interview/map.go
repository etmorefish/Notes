// 米哈游面试题
package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(map[int]int, 10)
	for i := 1; i <= 10; i++ {
		data[i] = i
	}
	for key, val := range data {
		/*
			go func() {
				fmt.Println("key: ", key, "val: ", val)
			}()
			打印出来的值会重复，应该强制刷新goroutine 栈中的值
		*/
		go func(key, val int) {
			fmt.Println("key: ", key, "val: ", val)
		}(key, val)
	}
	time.Sleep(time.Second * 5)
}
