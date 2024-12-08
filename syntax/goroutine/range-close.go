package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// close()后为false
	for i := range c {
		fmt.Println(i)
	}
	// 关闭通道后不能在接收数据
	// c <- 0
}
