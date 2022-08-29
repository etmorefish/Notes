package main

import "fmt"

func func1() func(int) int {
	sum := 0
	return func(val int) int {
		sum += val
		return sum
	}
}
func printFunc1() {
	sumFunc := func1()
	fmt.Println(sumFunc(1))
	fmt.Println(sumFunc(1))
}
func main() {
	printFunc1()
}
