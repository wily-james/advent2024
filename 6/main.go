package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var dirs [4][2]int = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type Iter struct {
	i   int
	j   int
	dir int

	lines [][]byte
}

func (it *Iter) Next() bool {
	r, c := len(it.lines), len(it.lines[0])

	dir := dirs[it.dir]
	ni, nj := it.i+dir[0], it.j+dir[1]
	if ni >= 0 && ni < r && nj >= 0 && nj < c {
		if it.lines[ni][nj] != '.' && it.lines[ni][nj] != '^' {
			it.dir = (it.dir + 1) % len(dirs)
			return true
		}
	}

	it.i, it.j = ni, nj
	if it.i < 0 || it.i >= r || it.j < 0 || it.j >= c {
		return false
	}
	return true
}

func main() {
	input := flag.String("i", "input", "input file name")
	partTwo := flag.Bool("2", false, "run part two")
	flag.Parse()

	file, err := os.Open(*input)
	check(err)
	defer file.Close()

	lines := make([][]byte, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		lines = append(lines, make([]byte, len(bytes)))
		copy(lines[len(lines)-1], bytes)
	}

	r, c := len(lines), len(lines[0])
	si, sj := -1, -1
	for i := range r {
		for j := range c {
			if lines[i][j] == '^' {
				si, sj = i, j
			}
		}
	}
	if si == -1 || sj == -1 {
		panic("No start position!")
	}
	iter := Iter{si, sj, 0, lines}

	if *partTwo {
		fmt.Println(countObstructionPoints(lines, iter))
	} else {
		fmt.Println(countVisited(lines, iter))
	}
}

func countObstructionPoints(lines [][]byte, iter Iter) int {
	r, c := len(lines), len(lines[0])

	seen := make([][]bool, r)
	for i := range r {
		seen[i] = make([]bool, c)
	}

	count := 0
	for {
		dir := dirs[iter.dir]
		ni, nj := iter.i+dir[0], iter.j+dir[1]
		if ni < 0 || ni >= r || nj < 0 || nj >= c {
			if !iter.Next() {
				break
			}
			continue
		}

		if lines[ni][nj] != '.' || seen[ni][nj] {
			if !iter.Next() {
				break
			}
			continue
		}
		seen[ni][nj] = true

		lines[ni][nj] = '#'
		if isCyclic(iter) {
			count += 1
		}
		lines[ni][nj] = '.'

		if !iter.Next() {
			break
		}
	}

	return count
}

func isCyclic(iter Iter) bool {
	slow, fast := iter, iter
	for {
		if !fast.Next() || !fast.Next() {
			return false
		}
		slow.Next()

		if slow.i == fast.i && slow.j == fast.j && slow.dir == fast.dir {
			return true
		}
	}
}

func countVisited(lines [][]byte, iter Iter) int {
	r, c := len(lines), len(lines[0])

	seen := make([][]bool, r)
	for i := range r {
		seen[i] = make([]bool, c)
	}
	seen[iter.i][iter.j] = true

	for iter.Next() {
		seen[iter.i][iter.j] = true
	}

	count := 0
	for i := range r {
		for j := range c {
			if seen[i][j] {
				count += 1
			}
		}
	}
	return count
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
