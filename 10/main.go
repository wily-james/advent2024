package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	input := flag.String("i", "input", "input file name")
	flag.Parse()

	file, err := os.Open(*input)
	check(err)
	defer file.Close()

	lines := make([][]byte, 0, 10)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		line := make([]byte, len(bytes))
		copy(line, bytes)
		lines = append(lines, line)
	}

	scoreSum, ratingSum := sumScores(lines)
	fmt.Printf("Score: %d\n", scoreSum)
	fmt.Printf("Rating: %d\n", ratingSum)
}

func sumScores(lines [][]byte) (int, int) {
	r, c := len(lines), len(lines[0])
	reachable := make([][]bool, r)
	for i := range r {
		reachable[i] = make([]bool, c)
	}

	scoreSum := 0
	ratingSum := 0
	for i := range r {
		for j := range c {
			ratingSum += score(lines, i, j, r, c, 0, reachable)

			for i := range r {
				for j := range c {
					if reachable[i][j] {
						scoreSum++
					}
					reachable[i][j] = false
				}
			}
		}
	}
	return scoreSum, ratingSum
}

func score(lines [][]byte, i, j, r, c, step int, reachable [][]bool) int {
	if i < 0 || i >= r || j < 0 || j >= c || lines[i][j] != ('0'+byte(step)) {
		return 0
	}

	if step == 9 {
		reachable[i][j] = true
		return 1
	}

	sum := 0
	for _, dir := range [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		sum += score(lines, i+dir[0], j+dir[1], r, c, step+1, reachable)
	}
	return sum
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
