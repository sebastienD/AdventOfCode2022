package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	sum := computeFile()
	fmt.Printf("La somme vaut %d\n", sum)
}

func computeFile() int64 {
	filename := "input.txt"
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	current := newRes("")
	arbo := make(map[string]resource)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$") {
			switch line[2:4] {
			case "ls":
				continue
			case "cd":
				if current.path != "" {
					arbo[current.path] = current
					//fmt.Println("arbo", arbo)
				}
				var newPath string
				fmt.Sscanf(line, "$ cd %s", &newPath)
				p := path.Join(current.path, newPath)
				//fmt.Println("capture", newPath, "calc", p)
				c, exists := arbo[p]
				if exists {
					current = c
				} else {
					current = newRes(p)
				}
			}
		} else {
			if strings.HasPrefix(line, "dir") {
				var dir string
				fmt.Sscanf(line, "dir %s", &dir)
				//fmt.Println(dir)
				current.dirs = append(current.dirs, path.Join(current.path, dir))
			} else {
				var size int
				var name string
				fmt.Sscanf(line, "%d %s", &size, &name)
				//fmt.Println(size, " and ", name)
				current.files = append(current.files, size)
			}
		}
	}
	arbo[current.path] = current
	//fmt.Println("arbo", arbo)

	max := int64(100000)
	_ = calc(arbo, arbo["/"], max)
	var total int64
	for _, v := range arbo {
		if v.total <= max {
			total += v.total
		}
	}
	return total
}

func calc(arbo map[string]resource, res resource, max int64) int64 {
	if res.total == -1 {
		res.total = 0
		for _, dir := range res.dirs {
			c := calc(arbo, arbo[dir], max)
			res.total += c
		}
		for _, v := range res.files {
			res.total += int64(v)
		}
	}
	arbo[res.path] = res
	//fmt.Println(res.path, res.total)
	return res.total
}

type resource struct {
	path  string
	dirs  []string
	files []int
	total int64
}

func newRes(p string) resource {
	return resource{path: p, total: -1}
}
