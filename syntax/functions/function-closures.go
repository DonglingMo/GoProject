package main

import "fmt"

func adder() func(int) int {
	// 每次调用adder都会有一个新的sum， pos和neg各对应一个
	sum := 0
	// 返回匿名函数闭包引用的sum
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			// 不同的算
			pos(i),
			neg(-2*i),
		)
	}
}
