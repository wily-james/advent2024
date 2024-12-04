package main

import (
	"bufio"
	"flag"
	"os"
)

const target string = "XMAS"

type Dir int

const (
	Search Dir = iota
	Up
	Down
	Left
	Right
	UpRight
	UpLeft
	DownRight
	DownLeft
)

func traverse(lines [][]byte, i, j, r, c, depth int, dir Dir) int {
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

	saved := lines[i][j]
	lines[i][j] = '\n'
	sum := 0

	switch dir {
	case Search:
		sum += traverse(lines, i+1, j, r, c, depth+1, Up)
		sum += traverse(lines, i-1, j, r, c, depth+1, Down)
		sum += traverse(lines, i+1, j+1, r, c, depth+1, Left)
		sum += traverse(lines, i-1, j-1, r, c, depth+1, Right)
		sum += traverse(lines, i+1, j-1, r, c, depth+1, UpRight)
		sum += traverse(lines, i-1, j+1, r, c, depth+1, UpLeft)
		sum += traverse(lines, i, j+1, r, c, depth+1, DownRight)
		sum += traverse(lines, i, j-1, r, c, depth+1, DownLeft)

	case Up:
		sum += traverse(lines, i+1, j, r, c, depth+1, Up)

	case Down:
		sum += traverse(lines, i-1, j, r, c, depth+1, Down)

	case Left:
		sum += traverse(lines, i+1, j+1, r, c, depth+1, Left)

	case Right:
		sum += traverse(lines, i-1, j-1, r, c, depth+1, Right)

	case UpRight:
		sum += traverse(lines, i+1, j-1, r, c, depth+1, UpRight)

	case UpLeft:
		sum += traverse(lines, i-1, j+1, r, c, depth+1, UpLeft)

	case DownRight:
		sum += traverse(lines, i, j+1, r, c, depth+1, DownRight)

	case DownLeft:
		sum += traverse(lines, i, j-1, r, c, depth+1, DownLeft)
	}
	lines[i][j] = saved
	return sum
}

func search(lines [][]byte, i, j, r, c int) int {
	return traverse(lines, i, j, r, c, 0, Search)
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
