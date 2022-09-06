package main

import (
	"fmt"
	"sync"
	"time"
)

/*
使用三个协程，每秒钟打印cat、dog、fish
顺序不能变换
无限循环即可
*/

func cat(fishCH, catCH chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			fmt.Println("cat")
			catCH <- struct{}{} // 写入值，告知其它 “cat 已打印”
			<-fishCH            // 从fish中读，会阻塞住
		}
	}()
}

func dog(catCH, dogCH chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			<-catCH
			fmt.Println("dog")
			dogCH <- struct{}{}
		}
		// wg.Done()
	}()

}

func fish(dogCH, fishCH chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			<-dogCH
			fmt.Println("fish")
			time.Sleep(time.Second)
			fishCH <- struct{}{}
		}
		// wg.Done()
	}()

}

func main() {
	catCH := make(chan struct{})
	dogCH := make(chan struct{})
	fishCH := make(chan struct{})
	wg := sync.WaitGroup{}
	
	cat(fishCH, catCH, &wg)
	dog(catCH, dogCH, &wg)
	fish(dogCH, fishCH, &wg)
	wg.Wait()

}
