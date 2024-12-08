package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// 参数列表类型简写
func add2(x, y int) int {
	return x + y
}

// 多值返回
func swap(x, y string) (string, string) {
	return y, x
}

// 返回值名称
func split(sum int) (x, y int) {
	x = sum / 2
	y = sum - x
	return
}

func main() {
	fmt.Println(add(3, 4)) // Output: 7

	a, b := swap("a", "b")
	fmt.Println(a, b) // Output: b a

	fmt.Println(split(17)) // Output: 8 9
}
