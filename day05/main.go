package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	stacks := build(false)
	tops := make([]string, len(stacks))
	for k, v := range stacks {
		index := len(v) - 1
		tops[Atoi(k)-1] = v[index]
	}
	fmt.Printf("Les tops sont %s\n", strings.Join(tops, ""))
	stacks = build(true)
	tops = make([]string, len(stacks))
	for k, v := range stacks {
		index := len(v) - 1
		tops[Atoi(k)-1] = v[index]
	}
	fmt.Printf("Les tops sont %s\n", strings.Join(tops, ""))
}

func build(conserveOrder bool) map[string][]string {
	filename := "input.txt"
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	stacks := map[string][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "    ", "_")
		line = strings.ReplaceAll(line, "[", "")
		line = strings.ReplaceAll(line, "]", "")
		line = strings.ReplaceAll(line, " ", "")
		if strings.HasPrefix(line, "1") {
			break
		}
		l := strings.Split(line, "")
		for i, v := range l {
			if v == "_" {
				continue
			}
			stack := stacks[strconv.Itoa(i+1)]
			stack = append(stack, v)
			stacks[strconv.Itoa(i+1)] = stack
		}
	}
	fmt.Println(stacks)
	for k := range stacks {
		reverse(stacks[k])
	}
	fmt.Println(stacks)
	scanner.Scan()
	reg := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		res := reg.FindStringSubmatch(line)
		num := Atoi(res[1])
		last := len(stacks[res[2]]) - num
		far := stacks[res[2]][last:]
		if !conserveOrder {
			reverse(far)
		}
		stacks[res[3]] = append(stacks[res[3]], far...)
		stacks[res[2]] = stacks[res[2]][:last]
	}

	return stacks
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
