package main

import "fmt"

func main() {
	var x *int = nil
	var y interface{} = x
	var z interface{} = nil

	fmt.Println(x == y)
	fmt.Println(x == nil)
	fmt.Println(y == nil)
	fmt.Println(z == y)
	fmt.Println(z == x)
}
