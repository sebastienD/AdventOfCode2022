package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type elf struct {
	sum   int
	count int
}

func main() {
	filename := "input.txt"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Read file %s: %v", filename, err)
	}
	defer f.Close()
	fileScanner := bufio.NewScanner(f)

	elves := []elf{}
	current := elf{}
	max := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			elves = append(elves, current)
			if current.sum > max {
				max = current.sum
			}
			current = elf{}
		}
		current.count++
		val, _ := strconv.Atoi(line)
		current.sum += val
	}
	fmt.Printf("Le max est %d\n", max)

	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].sum > elves[j].sum
	})
	top3 := elves[0].sum + elves[1].sum + elves[2].sum
	fmt.Printf("La somme du top 3 est de %d\n", top3)
}
