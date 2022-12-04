package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := computeFile(fullyIncludeCount)
	fmt.Printf("Le nombre de pairs incluses est %d\n", sum)
	sum = computeFile(overlapCount)
	fmt.Printf("Le nombre d'overlaps est %d\n", sum)
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

func fullyIncludeCount(line string) int {
	interS := strings.Split(line, ",")
	intervalles := [][]int{}
	for _, v := range interS {
		bornes := strings.Split(v, "-")
		intervalles = append(intervalles, []int{Atoi(bornes[0]), Atoi(bornes[1])})
	}

	if intervalles[0][0] == intervalles[1][0] || intervalles[0][1] == intervalles[1][1] {
		return 1
	}

	if intervalles[0][0] < intervalles[1][0] {
		return Btoi(intervalles[1][1] <= intervalles[0][1])
	}

	return Btoi(intervalles[0][1] <= intervalles[1][1])
}

func overlapCount(line string) int {
	interS := strings.Split(line, ",")
	intervalles := [][]int{}
	for _, v := range interS {
		bornes := strings.Split(v, "-")
		intervalles = append(intervalles, []int{Atoi(bornes[0]), Atoi(bornes[1])})
	}

	if intervalles[0][0] == intervalles[1][0] || intervalles[0][1] == intervalles[1][1] {
		return 1
	}

	if intervalles[0][0] < intervalles[1][0] {
		return Btoi(intervalles[1][0] <= intervalles[0][1])
	}

	return Btoi(intervalles[1][1] >= intervalles[0][0])
}

func Atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
