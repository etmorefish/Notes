package main

import (
	"./lib1"
	_ "./lib2"
	// _ "./lib2"
)

/*
import _ “fmt”
给fmt包起⼀一个别名，匿匿名， ⽆无法使⽤用当前包的⽅方法，但是
会执⾏行行当前的包内部的init()⽅方法

import aa “fmt”
给fmt包起⼀一个别名，aa， aa.Println()来直接调⽤用。

import . “fmt”
将当前fmt包中的全部⽅方法，导⼊入到当前本包的作⽤用中，fmt包中
的全部的⽅方法可以直接使⽤用API来调⽤用，不不需要fmt.API来调⽤用
*/

func main() {
	lib1.Lib1Test()

	// lib2.Lib2Test()
	// mylib2.Lib2Test()
	//Lib2Test()
}
