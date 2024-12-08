package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 库函数
	fmt.Println(strings.Join(os.Args[1:], " "))
}
