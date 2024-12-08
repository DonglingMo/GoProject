package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				_, err := fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				if err != nil {
					return
				}
				continue
			}
			countLines(f, counts)
			err = f.Close()
			if err != nil {
				return
			}
		}
	}
	for line, i := range counts {
		if i > 1 {
			fmt.Printf("%s: %d\n", line, i)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
