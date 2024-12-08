package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// map
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, i := range counts {
		if i > 1 {
			fmt.Printf("%s: %d\n", line, i)
		}
	}
}
