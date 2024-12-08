package main

import "fmt"

func main() {
	var s []int
	printSlice3(s)

	// 可在空切片上追加
	s = append(s, 0)
	printSlice3(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice3(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice3(s)
}

func printSlice3(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
