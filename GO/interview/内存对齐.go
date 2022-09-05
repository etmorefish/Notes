package main

import (
	"fmt"
	"unsafe"
)

type DemoA struct {
	A int8
	B int64
	C int16
}

type DemoB struct {
	A int8
	C int16
	B int64
}

func main() {
	fmt.Println(unsafe.Sizeof(DemoA{}))
	fmt.Println(unsafe.Sizeof(DemoB{}))

}
/*
32位机器，对齐4字节
64位机器，对齐8字节

int8 是1个字节，填充7字节
int16 是2个字节，填充6个字节
int64 是8个字节

*/