package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// 切片
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
