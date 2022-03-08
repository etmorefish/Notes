package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

// 在一个Goroutinne 重复上锁会发生panic

func main() {
	mu.Lock()
	go func() {
		mu.Lock()
		fmt.Println("ok")
		defer mu.Unlock()
	}()
	mu.Unlock()
	time.Sleep(time.Second)
}
