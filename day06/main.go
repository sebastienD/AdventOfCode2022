package main

import (
	"fmt"
	"os"
)

func main() {
	search(4)
	search(14)
}

func search(distinct int) {
	s, _ := os.ReadFile("input.txt")
	index := distinct
	for start, end := 0, distinct; end < len(s); start, end = start+1, end+1 {
		count := map[byte]int{}
		for i := start; i < end; i++ {
			count[s[i]]++
		}
		if len(count) == distinct {
			break
		}
		index++
	}
	fmt.Printf("L'index est %d\n", index)
}
