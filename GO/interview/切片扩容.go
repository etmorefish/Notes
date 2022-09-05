package main

import "fmt"

// 当原slice容量(oldcap)小于256的时候，
// 新slice(newcap)容量为原来的2倍；
// 原slice容量超过256，新slice容量newcap = oldcap+(oldcap+3*256)/4
func main() {
	var arr1 []int64
	for i := 0; i < 513; i++ {
		arr1 = append(arr1, int64(i))
	}
	fmt.Printf("len= %d, cap = %d \n", len(arr1), cap(arr1))
}
