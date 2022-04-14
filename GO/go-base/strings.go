package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello,世界"

	fmt.Println(len(s))
	fmt.Println(len([]byte(s)))
	fmt.Println(utf8.RuneCountInString(s))
}
