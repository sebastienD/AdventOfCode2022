package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	sum := computeFile(prio)
	fmt.Printf("La somme des priorités est %d\n", sum)
	sum = prioBy3(partitionBy3())
	fmt.Printf("La somme des priorités est %d\n", sum)
}

func computeFile(apply func(line string) int) int {
	filename := "input.txt"
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var total int
	for scanner.Scan() {
		v := apply(scanner.Text())
		total += v
	}
	return total
}

func partitionBy3() [][]string {
	filename := "input.txt"
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var pack [][]string
	var current []string
	var lineNum int
	for scanner.Scan() {
		if lineNum%3 == 0 && lineNum != 0 {
			pack = append(pack, current)
			current = []string{}
		}
		current = append(current, scanner.Text())
		lineNum++
	}
	pack = append(pack, current)
	return pack

}

func prioBy3(pack [][]string) int {
	var total int
	for _, v := range pack {
		r := searchRuneFor3(v)
		total += getPrio(r)
	}
	return total
}

func prio(l string) int {
	r := searchRune(l)
	return getPrio(r)
}

func searchRuneFor3(group []string) rune {
	l1 := []rune(group[0])
	l2 := []rune(group[1])
	l3 := []rune(group[2])
	for _, v := range l1 {
		for _, k := range l2 {
			if k == v {
				for _, p := range l3 {
					if k == p {
						return k
					}
				}
			}
		}
	}
	return 0
}

func searchRune(l string) rune {
	bags := []rune(l)
	s := len(bags) / 2
	first := bags[:s]
	sort.Slice(first, func(i, j int) bool {
		return first[i] < first[j]
	})
	second := bags[s:]
	sort.Slice(second, func(i, j int) bool {
		return second[i] < second[j]
	})
	for _, v := range first {
		for _, k := range second {
			if v == k {
				return k
			}
			if k > v {
				break
			}
		}
	}
	return 0
}

func getPrio(r rune) int {
	if r > 96 {
		return int(r) - 96
	}
	return int(r) - 38
}
