package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	score := first()
	fmt.Printf("Le score total est %d\n", score)
	score = second()
	fmt.Printf("Le score total est %d\n", score)
}

func first() int {
	filename := "input.txt"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Read file %s: %v", filename, err)
	}
	defer f.Close()
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	var total int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		game := strings.Split(line, " ")
		total += score(game[0], game[1])
	}
	return total
}

func second() int {
	filename := "input.txt"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Read file %s: %v", filename, err)
	}
	defer f.Close()
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	var total int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		game := strings.Split(line, " ")
		s := secondScore(game[0], game[1])
		total += s
	}
	return total
}

func secondScore(op, res string) int {
	total := winScore(res)
	switch op + res { // X perte
	case "AX":
		total += pscore("Z")
	case "AY":
		total += pscore("X")
	case "AZ":
		total += pscore("Y")
	case "BX":
		total += pscore("X")
	case "BY":
		total += pscore("Y")
	case "BZ":
		total += pscore("Z")
	case "CX":
		total += pscore("Y")
	case "CY":
		total += pscore("Z")
	case "CZ":
		total += pscore("X")
	}
	return total
}

func score(op, my string) int {
	total := pscore(my)
	switch op + my {
	case "AX":
		total += 3
	case "AY":
		total += 6
	case "AZ":
		total += 0
	case "BX":
		total += 0
	case "BY":
		total += 3
	case "BZ":
		total += 6
	case "CX":
		total += 6
	case "CY":
		total += 0
	case "CZ":
		total += 3
	}
	return total
}

func pscore(my string) int {
	switch my {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	}
	return 0
}

func winScore(result string) int {
	switch result {
	case "X":
		return 0
	case "Y":
		return 3
	case "Z":
		return 6
	}
	return 0
}
