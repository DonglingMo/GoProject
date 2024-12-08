package main

import "fmt"

// 常量 字符常量， 字符串，bool 常量，数值常量首字母大写
const Pi = 3.14

func main() {
	const World = "Hello, 世界"
	fmt.Println(Pi, World) // 3.14 Hello, 世界
	const Truth = true
	fmt.Println(Truth) // true
}
