package main

import (
	"fmt"
	"sync"
	"time"
)

//dog cat fish 按这个输出一百句   第一题

var wg sync.WaitGroup

func main() {
	// counter use to count the output
	var counter int = 100
	dogch := make(chan struct{}, 1)
	catch := make(chan struct{}, 1)
	fishch := make(chan struct{}, 1)
	wg.Add(1)
	// let run the three goroutine

	// the origin of the goroutine to set dogch <- struct {}{}
	dogch <- struct{}{}

	go dog(dogch, counter, catch)
	go cat(catch, counter, fishch)
	go fish(fishch, counter, dogch)

	wg.Wait()
}

func dog(dogch chan struct{}, counter int, catch chan struct{}) {
	dogtimes := 0
	for dogtimes < counter {
		select {
		case <-dogch:
			fmt.Println("dog")
			dogtimes++
			// let cat run
			catch <- struct{}{}
		}
	}
	wg.Done()

}

func cat(catch chan struct{}, counter int, fishch chan struct{}) {
	cattimes := 0
	for cattimes < counter {
		select {
		case <-catch:
			fmt.Println("cat")
			cattimes++
			//let fish run
			fishch <- struct{}{}
		}
	}
	wg.Done()
}

func fish(fishch chan struct{}, counter int, dogch chan struct{}) {
	fishtimes := 0
	for fishtimes < counter {
		select {
		case <-fishch:
			fmt.Println("fish")
			fishtimes++
			// let dog run
			time.Sleep(time.Second)
			dogch <- struct{}{}
		}
	}
	wg.Done()
}

//题解：
// 分别使用三个goroutine代表dog cat fish，由于携程执行的无顺序性，但是题目要求顺序循环输出，
// 所以，我们使用通道来实现阻塞，通过传值到通道中来实现顺序执行
// 首先定义了三个有缓冲通道 ，并设置了循环的次数counter
// counter use to count the output
//	var counter int = 3
//	dogch := make(chan struct{}, 1)
//	catch := make(chan struct{}, 1)
//	fishch := make(chan struct{}, 1)

// 然后对于每一个goroutine来说设置一个 for循环来执行counter次，里面采用for-select-case
// 如果该通道没值 则会阻塞在这个地方直到有值传进来。

//  // the origin of the goroutine to set dogch <- struct {}{}
//  dogch <- struct{}{} 这一句话是三个goroutin启动的触发条件。
//  第一个goroutine有值之后，执行第一个输出 dog，然后向catch里面传值，这个时候被阻塞的第一个goroutine//
//  得以激活，一次类推，直到把循环走完，每个goroutine都done了，程序结束
