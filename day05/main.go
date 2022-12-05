package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	stacks := extract()
	var tops string
	for _, v := range stacks {
		index := len(v) - 1
		tops += v[index]
	}
	fmt.Printf("Les tops sont %s\n", tops)
}

func extract() [][]string {
	filename := "input.txt"
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	stacks := [][]string{}
	for scanner.Scan() {
		//re := regexp.MustCompile(`(    )|([A-Z]+ ){8}(   )|([A-Z]+)`)
		// idx 1 5 9
		line := strings.Split(scanner.Text(), "")
		if line[1] == "1" {
			break
		}
		fmt.Println(line[0], line[1])
		stacks[0] = append(stacks[0], line[1], line[5], line[9], line[13], line[17], line[21], line[24], line[29], line[33])
		break
	}
	return stacks
}
