package main

import (
	"fmt"
	"strings"
)

// 只增大不减小的滑动窗口
// 窗口内没有重复字符：此时判断i+1与end的关系，超过表示遍历到窗口之外了，增大窗口大小
// 窗口内出现重复字符：此时两个指针都增大index+1，滑动窗口位置到重复字符的后一位
// 遍历结束，返回end-start，窗口大小

/*
	func lengthOfLongestSubstring(s string) int {
		start := 0
		end := 0
		for i := 0; i < len(s); i++ {
			index := strings.Index(s[start:i], string(s[i]))
			if index == -1 && (i+1) > end {
				end = i + 1
			} else {
				start += index + 1
				end += index + 1
			}
		}
		return end - start
	}
*/

func lengthOfLongestSubstring(s string) int {
	start := 0
	end := 0
	var ret int
	for i := range s {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			end++
		} else {
			start += index + 1
			end++
		}
		ret = max(len(s[start:end]), ret)
	}
	return ret
}
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}
