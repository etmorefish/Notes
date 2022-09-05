package main

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁、不可重复锁 在一个Goroutine下，重复上锁会引发panic
var mu sync.Mutex

func main() {
	mu.Lock()
	go func() {
		mu.Lock()
		fmt.Println("Running")
		defer mu.Unlock()
	}()
	mu.Unlock()
	time.Sleep(time.Second)
	return
}
