package main

import "fmt"

func main() {
	// defer fmt.Printf("main func2 %d\n", func2())
	defer fmt.Printf("main func3 %d\n", func3())

}

func func2() (sum int) {
	sumA := 100
	sumB := 100
	sum = sumA + sumB
	defer func() {
		fmt.Printf("func2 first %d\n", sum)
	}()
	defer fmt.Printf("func2 second %d\n", sum)
	return sum * 10
}

/*
func2 second 200
func2 first 2000
main 2000
*/

func func3() int {
	sumA := 100
	sumB := 100
	sum := sumA + sumB
	defer func() {
		fmt.Printf("func2 first %d\n", sum)
	}()
	defer fmt.Printf("func2 second %d\n", sum)
	return sum * 10
}

/*
func2 second 200
func2 first 200
main func3 2000
*/
