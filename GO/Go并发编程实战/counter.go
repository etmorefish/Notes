package main

import (
	"fmt"
	"sync"
)

func main() {
	//互斥锁保护计数器
	var mu sync.Mutex
	//计数器的值
	var count = 0
	//使用WaitGroup 等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			//对变量执行10w次加1
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	//等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}
