package main

import "fmt"

// 1.
// func main() {
// 	s := make([]int, 5)
// 	s = append(s, 1, 2, 3)
// 	fmt.Println(s)
// }
// [0 0 0 0 0 1 2 3]

// 2.
// func main() {
// 	s := make([]int, 0)
// 	s = append(s, 1, 2, 3, 4)
// 	fmt.Println(s)
// }
// [1 2 3 4]

// 参考解析：这道题考的是使用 append 向 slice 添加元素，第一段代码常见的错误是 [1 2 3]，需要注意。

// 3.下面这段代码有什么缺陷
// func funcMui(x,y int)(sum int,error){
// 	return x+y,nil
// }
func funcMui(x, y int) (sum int, err error) {
	return x + y, nil
}
func main() {
	res, err := funcMui(3, 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

// 第二个返回值没有命名。
// 参考解析：
// 在函数有多个返回值时，只要有一个返回值有命名，其他的也必须命名。
// 如果有多个返回值必须加上括号()；如果只有一个返回值且命名也必须加上括号()。
// 这里的第一个返回值有命名 sum，第二个没有命名，所以错误


/*
3.new() 与 make() 的区别
参考答案：

new(T) 和 make(T,args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同。

new(T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T的值。
换句话说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组、结构体等。

make(T,args) 返回初始化之后的 T 类型的值，这个值并不是 T 类型的零值，也不是指针 *T，
是经过初始化之后的 T 的引用。make() 只适用于 slice、map 和 channel.
*/