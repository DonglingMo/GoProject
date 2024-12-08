package main

import "fmt"

// 匿名函数，包级函数
func square() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := square()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 4
	fmt.Println(f()) // 9
}
