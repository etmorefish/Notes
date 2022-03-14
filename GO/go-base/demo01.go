package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/home/lei/t.sh")
	defer f.Close()
	for {
		n, err := f.Read(buf)
		if err != nil {
			fmt.Println("err: ", err)
		}
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}

	const s = 1 << 8
	fmt.Println(s)

	s1 := "Hello, world!"

	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(strings.Split(s1, ","))
	fmt.Println(strings.Contains(s1, "o"))
	fmt.Println(strings.HasPrefix(s1, "h"))
	fmt.Println(strings.Index(s1, "o"))
	fmt.Println(strings.Join(strings.Split(s1, ","), "---"))

	traversalString()
	changeString()
	sqrtDemo()
}

// 遍历字符串
func traversalString() {
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

//修改字符串
func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}

//类型转换
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
