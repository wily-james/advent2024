package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(l int) int {
	if l > 0 {
		return l
	}
	return -l
}

func isSafe(levels []int) bool {
	if len(levels) > 1 {
		if levels[0] == levels[1] {
			return false
		}

		increasing := levels[1] > levels[0]
		for i := 1; i < len(levels); i++ {
			if increasing && levels[i] <= levels[i-1] {
				return false
			}
			if !increasing && levels[i] >= levels[i-1] {
				return false
			}

			if abs(levels[i]-levels[i-1]) < 1 || abs(levels[i]-levels[i-1]) > 3 {
				return false
			}

			if i+1 < len(levels) && (abs(levels[i]-levels[i+1]) < 1 || abs(levels[i]-levels[i+1]) > 3) {
				return false
			}
		}
	}

	return true
}

func part1(reports [][]int) int {
	safe := 0
	for _, levels := range reports {
		if isSafe(levels) {
			safe += 1
		}
	}

	return safe
}

func part2(reports [][]int) int {
	safe := 0
	for _, levels := range reports {
		if isSafe(levels) {
			safe += 1
		} else {
			for i := range levels {
				newLevels := make([]int, 0, len(levels)-1)
				for j := range levels {
					if i == j {
						continue
					}
					newLevels = append(newLevels, levels[j])
				}
				if isSafe(newLevels) {
					safe += 1
					break
				}
			}
		}
	}

	return safe
}

func main() {
	part := flag.Int("p", 1, "enter 1 for part1, or 2 for part 2")
	fileName := flag.String("i", "input", "input filename. Default is 'input'")
	flag.Parse()

	file, err := os.Open(*fileName)
	check(err)
	defer file.Close()

	reports := make([][]int, 0, 100)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reports = append(reports, make([]int, 0))
		for _, level := range strings.Fields(scanner.Text()) {
			r, err := strconv.Atoi(level)
			check(err)
			reports[len(reports)-1] = append(reports[len(reports)-1], r)
		}
	}

	if *part == 1 {
		safe := part1(reports)
		fmt.Println(safe)
	} else if *part == 2 {
		safe := part2(reports)
		fmt.Println(safe)
	} else {
		check(fmt.Errorf("Choose part 1 or part 2"))
	}
}

func check(e error) {
	if e != nil {
		fmt.Printf("Error: %v", e)
		os.Exit(1)
	}
}
