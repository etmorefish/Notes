package main

import (
	"fmt"
	"sync"
)

func main() {
	letter, number := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				// fmt.Print(i)
				// i++
				letter <- true
			}
		}
	}()
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-letter:
				if i >= 'Z' {
					wg.Done()
					return
				}

				fmt.Print(string(i))
				i++
				// fmt.Print(string(i))
				// i++
				number <- true
			}

		}
	}(&wg)
	number <- true
	wg.Wait()
}
