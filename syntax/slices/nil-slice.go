package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	// 空切片
	if s == nil {
		fmt.Println("nil!")
	}
}
