package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			if err != nil {
				return
			}
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%s: %d\n", line, count)
		}
	}
}
