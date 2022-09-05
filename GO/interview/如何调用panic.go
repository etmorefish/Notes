package main

import (
	"fmt"
	"time"
)

func proc() {
	panic("ok!")
}

/*
要求每秒调用一次proc函数
*/
func main() {
	go func() {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-t.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					proc()
				}()
			}
		}
	}()

	// select {}
	time.Sleep(time.Second * 10)
}
