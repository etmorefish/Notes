package main

import "fmt"

func func3() int {
	val := 10
	defer func() {
		val += 1
		fmt.Println(val)
	}()
	return val
}

func printFunc3() {
	fmt.Println(func3())
}
func main() {
	printFunc3()
}
