package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	// cap == len
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	// len = 0 , cap = 底层数组长度
	printSlice(s)

	// 扩展其长度
	s = s[:4]
	// len = 4 cap = 6
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	// 从index == 2开始到底层数组的最后一位
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
