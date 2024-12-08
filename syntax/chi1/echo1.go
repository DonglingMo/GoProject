package main

import (
	"fmt"
	"os"
)

func main() {
	// 变量声明
	var s, sep string
	// os参数读取，循环
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
