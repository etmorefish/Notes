package main

import "fmt"

func changeValue(p int) {
	p = 10
}

func main() {
	a := 1
	changeValue(a)
	fmt.Println("a = ", a)
}

// 传值
