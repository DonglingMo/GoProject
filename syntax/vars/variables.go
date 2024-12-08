package main

import "fmt"

var c, python, java bool

// var声明函数或者package层级变量
// var 变量声明并初始化时，可以不用带声明类型，完成类型推断
func main() {
	var i int
	fmt.Println(i, c, python, java) // 0 false false false
}
