package main

import "fmt"

func main() {
	//defer fmt.Println("world")
	//
	//fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// 函数调用好入栈
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
