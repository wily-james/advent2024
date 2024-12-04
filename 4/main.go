package main

import (
	"bufio"
	"flag"
	"os"
)

const target string = "XMAS"

type Dir struct {
	i int
	j int
}

var AllDirs [8]Dir = [8]Dir{
	{0, 1},
	{0, -1},
	{1, 0},
	{1, 1},
	{1, -1},
	{-1, 0},
	{-1, 1},
	{-1, -1},
}

func searchDir(lines [][]byte, i, j, r, c, depth int, dir Dir) int {
	if depth == len(target) {
		return 0
	}
	if i >= r || i < 0 || j >= c || j < 0 {
		return 0
	}
	if lines[i][j] != target[depth] {
		return 0
	}
	if depth == len(target)-1 {
		return 1
	}

	return searchDir(lines, i+dir.i, j+dir.j, r, c, depth+1, dir)
}

func search(lines [][]byte, i, j, r, c int) int {
	if lines[i][j] != target[0] {
		return 0
	}

	sum := 0
	for _, dir := range AllDirs {
		sum += searchDir(lines, i+dir.i, j+dir.j, r, c, 1, dir)
	}
	return sum
}

func doPartTwo(lines [][]byte, r, c int) int {
	count := 0
	for i := range r - 2 {
		for j := range c - 2 {
			if isXmas(lines, i, j) {
				count += 1
			}
		}
	}
	return count
}

func isXmas(lines [][]byte, i, j int) bool {
	if lines[i+1][j+1] != 'A' {
		return false
	}

	return ((lines[i][j] == 'M' && lines[i+2][j+2] == 'S') || (lines[i][j] == 'S' && lines[i+2][j+2] == 'M')) && ((lines[i+2][j] == 'M' && lines[i][j+2] == 'S') || (lines[i+2][j] == 'S' && lines[i][j+2] == 'M'))
}

func main() {
	fileName := flag.String("i", "input", "input file name")
	partTwo := flag.Bool("2", false, "pass to run part two")
	flag.Parse()

	file, err := os.Open(*fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([][]byte, 0, 10)
	for scanner.Scan() {
		lines = append(lines, []byte(scanner.Text()))
	}

	r, c := len(lines), len(lines[0])

	if *partTwo {
		println(doPartTwo(lines, r, c))
	} else {
		count := 0
		for i := range r {
			for j := range c {
				count += search(lines, i, j, r, c)
			}
		}
		println(count)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
