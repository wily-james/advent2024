package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	input := flag.String("i", "input", "input file name")
	partTwo := flag.Bool("2", false, "run part two")
	flag.Parse()

	file, err := os.Open(*input)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([][]byte, 0, 10)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		line := make([]byte, len(bytes))
		copy(line, bytes)
		lines = append(lines, line)
	}

	if *partTwo {
		fmt.Println(antiNodePositions2(lines))
	} else {
		fmt.Println(antiNodePositions(lines))
	}
}

func antiNodePositions(lines [][]byte) int {
	r, c := len(lines), len(lines[0])
	m := make([][]bool, r)
	for i := range r {
		m[i] = make([]bool, c)
	}

	for i := range r {
		for j := range c {
			if lines[i][j] == '.' {
				continue
			}

			for i2 := range r {
				for j2 := range c {
					if i == i2 && j == j2 {
						continue
					}
					if lines[i][j] != lines[i2][j2] {
						continue
					}

					ai, aj := i-(i2-i), j-(j2-j)
					if ai >= 0 && ai < r && aj >= 0 && aj < c {
						m[ai][aj] = true
					}

					ai, aj = i2+(i2-i), j2+(j2-j)
					if ai >= 0 && ai < r && aj >= 0 && aj < c {
						m[ai][aj] = true
					}
				}
			}
		}
	}

	count := 0
	for i := range r {
		for j := range c {
			if m[i][j] {
				count++
			}
		}
	}
	return count
}

func antiNodePositions2(lines [][]byte) int {
	r, c := len(lines), len(lines[0])
	m := make([][]bool, r)
	for i := range r {
		m[i] = make([]bool, c)
	}

	for i := range r {
		for j := range c {
			if lines[i][j] == '.' {
				continue
			}

			for i2 := range r {
				for j2 := range c {
					if i == i2 && j == j2 {
						continue
					}
					if lines[i][j] != lines[i2][j2] {
						continue
					}

					di, dj := i2-i, j2-j
					for ai, aj := i, j; ai >= 0 && ai < r && aj >= 0 && aj < c; ai, aj = ai-di, aj-dj {
						m[ai][aj] = true
					}

					for ai, aj := i2, j2; ai >= 0 && ai < r && aj >= 0 && aj < c; ai, aj = ai+di, aj+dj {
						m[ai][aj] = true
					}
				}
			}
		}
	}

	count := 0
	for i := range r {
		for j := range c {
			if m[i][j] {
				count++
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
