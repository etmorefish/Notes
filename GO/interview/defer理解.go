package main

import "fmt"

func main() {
	m := 10
	defer fmt.Printf("first defer %d\n", +m)
	m = 100
	defer func() {
		fmt.Printf("second defer %d\n", m)
	}()
	m *= 10
	defer fmt.Printf("third defer %d\n", m)
	funcVal := func1()
	funcVal()
	m *= 10

}

func func1() func() {
	defer fmt.Println("before return")
	return func() {
		defer fmt.Println("in the return")
	}
}
