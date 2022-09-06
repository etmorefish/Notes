package main

import (
	"fmt"
	"net"
	"runtime/pprof"
	"sync"
)

/*
不会回收线程
*/
var threadProfile = pprof.Lookup("threadcreate")

func main() {
	fmt.Printf("协程执行前的线程数：%d\n", threadProfile.Count())
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				// 纯go 调用/etc/resolv.conf
				// cgo 调用c标准库 GODEBUG=netdns=cgo go run golang会不会回收线程.go
				_, err := net.LookupHost("www.baidu.com")
				if err != nil {
					return
				}
			}
		}()
	}
	wg.Wait()
	fmt.Printf("协程执行之后剩下的线程数：%d\n", threadProfile.Count())
}
