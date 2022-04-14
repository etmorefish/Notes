package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("str := \"c:\\pprof\\main.exe\"")

	buf := make([]byte, 1024)
	f, _ := os.Open("/home/lei/Documents/Notes/GO/go-base/arr01.go")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf)
	}



}
