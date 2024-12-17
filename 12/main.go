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

	lines2 := copyLines(lines)
	fmt.Println(totalPrice(lines))
	fmt.Println(bulkPrice(lines2))
}

func bulkPrice(lines [][]byte) int {
	if len(lines) == 0 || len(lines[0]) == 0 {
		return 0
	}

	seenSides := make([][]byte, 0, len(lines))
	for i := range len(lines) {
		seenSides = append(seenSides, make([]byte, len(lines[i])))
	}

	sum := 0
	r, c := len(lines), len(lines[0])
	for i := range r {
		for j := range c {
			if lines[i][j] == '\n' {
				continue
			}

			col := lines[i][j]
			sides := numSides(lines, i, j, r, c, col, seenSides)
			color(lines, i, j, r, c, ' ', col)
			for i := range r {
				for j := range c {
					seenSides[i][j] = 0
				}
			}

			area, _ := explore(lines, i, j, r, c, lines[i][j])
			color(lines, i, j, r, c, ' ', '\n')

			sum += area * sides
		}
	}
	return sum
}

func copyLines(lines [][]byte) [][]byte {
	lines2 := make([][]byte, 0, len(lines))
	for i := range len(lines) {
		line := make([]byte, len(lines[i]))
		copy(line, lines[i])
		lines2 = append(lines2, line)
	}
	return lines2
}

var dirs [][2]int = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func numSides(lines [][]byte, i, j, r, c int, b byte, seenSides [][]byte) int {
	lines[i][j] = ' '
	sides := 0
	for di, dir := range dirs {
		ni, nj := i+dir[0], j+dir[1]
		if ni < 0 || ni >= r || nj < 0 || nj >= c {
			sides += traceSide(lines, i, j, r, c, b, (di+1)%len(dirs), seenSides)
			continue
		}

		if lines[ni][nj] == ' ' {
			continue
		}

		if lines[ni][nj] != b {
			sides += traceSide(lines, i, j, r, c, b, (di+1)%len(dirs), seenSides)
			continue
		}

		sides += numSides(lines, ni, nj, r, c, b, seenSides)
	}

	return sides
}

func traceSide(lines [][]byte, si, sj, r, c int, b byte, sdi int, seenSides [][]byte) int {
	if (seenSides[si][sj] & (1 << ((sdi + len(dirs) - 1) % len(dirs)))) > 0 {
		return 0
	}

	sides := 0
	i, j, di := si, sj, sdi
	start := true
	for start || !(i == si && j == sj && di == sdi) {
		start = false
		dir := dirs[di]

		// Test cell to left to make sure we are still on wall
		ldi := (di + len(dirs) - 1) % len(dirs)
		ldir := dirs[ldi]
		li, lj := i+ldir[0], j+ldir[1]
		if li >= 0 && li < r && lj >= 0 && lj < c && (lines[li][lj] == b || lines[li][lj] == ' ') {
			// Turn left to stay on wall
			di = (di + len(dirs) - 1) % len(dirs)
			sides++
			i, j = li, lj

			// Mark wall to left as seen
			seenSides[i][j] |= (1 << ((di + len(dirs) - 1) % len(dirs)))
			continue
		}

		// Mark wall to left as seen
		seenSides[i][j] |= (1 << ldi)

		ni, nj := i+dir[0], j+dir[1]
		if ni < 0 || ni >= r || nj < 0 || nj >= c || (lines[ni][nj] != b && lines[ni][nj] != ' ') {
			di = (di + 1) % len(dirs) // turn right when hitting wall
			sides++
			continue
		}

		i, j = ni, nj
	}

	return sides
}

func totalPrice(lines [][]byte) int {
	if len(lines) == 0 || len(lines[0]) == 0 {
		return 0
	}

	sum := 0
	r, c := len(lines), len(lines[0])
	for i := range r {
		for j := range c {
			if lines[i][j] == '\n' {
				continue
			}

			area, perimeter := explore(lines, i, j, r, c, lines[i][j])
			color(lines, i, j, r, c, ' ', '\n')
			sum += area * perimeter
		}
	}
	return sum
}

func explore(lines [][]byte, i, j, r, c int, b byte) (int, int) {
	perim := 0
	area := 1
	lines[i][j] = ' '
	for _, dir := range dirs {
		ni, nj := i+dir[0], j+dir[1]
		if ni < 0 || ni >= r || nj < 0 || nj >= c {
			perim++
			continue
		}

		// Part of this region
		if lines[ni][nj] == ' ' {
			continue
		}

		if lines[ni][nj] != b {
			perim++
			continue
		}

		a, p := explore(lines, ni, nj, r, c, b)
		area += a
		perim += p
	}
	return area, perim
}

func color(lines [][]byte, i, j, r, c int, test, col byte) {
	lines[i][j] = col
	for _, dir := range dirs {
		ni, nj := i+dir[0], j+dir[1]
		if ni < 0 || ni >= r || nj < 0 || nj >= c {
			continue
		}

		if lines[ni][nj] != test {
			continue
		}

		color(lines, ni, nj, r, c, test, col)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
