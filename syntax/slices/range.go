package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i
	}
	// 忽略索引
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}
