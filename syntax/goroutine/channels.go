package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// 变量发送到信道
	c <- sum // 发送 sum 到 c
}

// 发送和接收端就绪前阻塞在，任意端准备好前阻塞，保证同步
func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// 创建一个通道，用于返回两个 goroutine 的返回值
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// 从信道接收赋值
	x, y := <-c, <-c // 从 c 接收

	fmt.Println(x, y, x+y)

	// 缓冲满时再发送数据会阻塞，缓冲区空时接收数据会阻塞
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
