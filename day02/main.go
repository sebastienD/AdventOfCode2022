package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var resScore = map[string]int{"X": 1, "Y": 2, "Z": 3}

func main() {
	score := computeFile(scoring)
	fmt.Printf("Le score total est %d\n", score)
	score = computeFile(secondScoring)
	fmt.Printf("Le score total est %d\n", score)
}

func computeFile(apply func(op, my string) int) int {
	filename := "input.txt"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Read file %s: %v", filename, err)
	}
	defer f.Close()
	fileScanner := bufio.NewScanner(f)

	var total int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		game := strings.Split(line, " ")
		total += apply(game[0], game[1])
	}
	return total
}

func secondScoring(op, res string) int {
	intRes := map[string]int{"X": 0, "Y": 3, "Z": 6}
	valRes := intRes[res]
	fScore := map[string]int{
		"AX": valRes + resScore["Z"],
		"AY": valRes + resScore["X"],
		"AZ": valRes + resScore["Y"],
		"BX": valRes + resScore["X"],
		"BY": valRes + resScore["Y"],
		"BZ": valRes + resScore["Z"],
		"CX": valRes + resScore["Y"],
		"CY": valRes + resScore["Z"],
		"CZ": valRes + resScore["X"],
	}
	return fScore[op+res]
}

func scoring(op, my string) int {
	rs := resScore[my]
	fScore := map[string]int{
		"AX": rs + 3,
		"AY": rs + 6,
		"AZ": rs + 0,
		"BX": rs + 0,
		"BY": rs + 3,
		"BZ": rs + 6,
		"CX": rs + 6,
		"CY": rs + 0,
		"CZ": rs + 3,
	}
	return fScore[op+my]
}
