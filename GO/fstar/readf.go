package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// var tf = []string{"/home/lei/Downloads/1011log/JB05", "/home/lei/Downloads/1011log/JB06"}
	var tf = []string{"/home/lei/Downloads/1011log/LUR1.go", "/home/lei/Downloads/1011log/LUR2.go"}

	file, err := os.OpenFile(tf[0], os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}
}


export https_proxy=http://192.168.1.10:7890 http_proxy=http://192.168.1.10:7890 all_proxy=socks5://192.168.1.10:7890