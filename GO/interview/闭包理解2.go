package main

import "fmt"

func func2() (val int) {
	val = 10
	defer func() {
		val += 1
	}()
	return val
}

func printFunc2() {
	fmt.Println(func2())
}
func main() {
	printFunc2()
}
