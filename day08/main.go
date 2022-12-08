package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	visibleTrees := first()
	fmt.Printf("Il y a %d arbres visibles\n", visibleTrees)
	score := second()
	fmt.Printf("The highest score is %d \n", score)
}

func first() int {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var ground [][]int
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), "")
		var l []int
		for _, v := range t {
			l = append(l, Atoi(v))
		}
		ground = append(ground, l)
	}
	//fmt.Println(ground)

	width := len(ground[0])
	height := len(ground)
	//fmt.Println(width, heigth)
	vt := width*2 + (height-2)*2

	for h := 1; h < height-1; h++ {
		for w := 1; w < width-1; w++ {
			//fmt.Println("point", ground[h][w])
			if isVisibleTop(ground, h, w) ||
				isVisibleRight(ground, h, w, width) ||
				isVisibleBottom(ground, h, w, height) ||
				isVisibleLeft(ground, h, w) {
				vt++
				//fmt.Println("point", ground[h][w], "visible")
			}
		}
	}

	return vt
}

func isVisibleTop(ground [][]int, h, w int) bool {
	for i := 1; i <= h; i++ {
		if ground[h][w] <= ground[h-i][w] {
			return false
		}
	}
	return true
}

func isVisibleRight(ground [][]int, h, w int, width int) bool {
	for i := 1; i <= (width - 1 - w); i++ {
		if ground[h][w] <= ground[h][w+i] {
			return false
		}
	}
	return true
}

func isVisibleBottom(ground [][]int, h, w int, height int) bool {
	for i := 1; i <= (height - 1 - h); i++ {
		if ground[h][w] <= ground[h+i][w] {
			return false
		}
	}
	return true
}

func isVisibleLeft(ground [][]int, h, w int) bool {
	for i := 1; i <= w; i++ {
		if ground[h][w] <= ground[h][w-i] {
			return false
		}
	}
	return true
}

func second() int {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var ground [][]int
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), "")
		var l []int
		for _, v := range t {
			l = append(l, Atoi(v))
		}
		ground = append(ground, l)
	}
	//fmt.Println(ground)

	width := len(ground[0])
	height := len(ground)
	//fmt.Println(width, heigth)

	var highest int
	for h := 1; h < height-1; h++ {
		for w := 1; w < width-1; w++ {
			score := scoreTop(ground, h, w) * scoreRight(ground, h, w, width) *
				scoreBottom(ground, h, w, height) * scoreLeft(ground, h, w)
			if score > highest {
				highest = score
			}
		}
	}

	return highest
}

func scoreTop(ground [][]int, h, w int) int {
	var score int
	for i := 1; i <= h; i++ {
		score++
		if ground[h][w] <= ground[h-i][w] {
			break
		}
	}
	return score
}

func scoreRight(ground [][]int, h, w int, width int) int {
	var score int
	for i := 1; i <= (width - 1 - w); i++ {
		score++
		if ground[h][w] <= ground[h][w+i] {
			break
		}
	}
	return score
}

func scoreBottom(ground [][]int, h, w int, height int) int {
	var score int
	for i := 1; i <= (height - 1 - h); i++ {
		score++
		if ground[h][w] <= ground[h+i][w] {
			break
		}
	}
	return score
}

func scoreLeft(ground [][]int, h, w int) int {
	var score int
	for i := 1; i <= w; i++ {
		score++
		if ground[h][w] <= ground[h][w-i] {
			break
		}
	}
	return score
}

func Atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
