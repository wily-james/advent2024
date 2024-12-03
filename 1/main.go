package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part2(lefts, rights []int) int {
	m := make(map[int]int, len(rights))
	for _, r := range rights {
		m[r] += 1
	}

	sim := 0
	for _, l := range lefts {
		sim += l * m[l]
	}
	return sim
}

func part1(lefts, rights []int) int {
	sort.Sort(sort.IntSlice(lefts))
	sort.Sort(sort.IntSlice(rights))

	dist := 0
	for i := range lefts {
		res := lefts[i] - rights[i]
		if res > 0 {
			dist += res
		} else {
			dist -= res
		}
	}

	return dist
}

func main() {
	part := flag.Int("p", 1, "enter 1 for part1, or 2 for part 2")
	fileName := flag.String("i", "input", "input filename. Default is 'input'")
	flag.Parse()

	file, err := os.Open(*fileName)
	check(err)
	defer file.Close()

	lefts := make([]int, 0, 100)
	rights := make([]int, 0, 100)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		left, err := strconv.Atoi(words[0])
		check(err)
		right, err := strconv.Atoi(words[1])
		check(err)

		lefts = append(lefts, left)
		rights = append(rights, right)
	}

	if *part == 1 {
		dist := part1(lefts, rights)
		fmt.Println(dist)
	} else if *part == 2 {
		similarity := part2(lefts, rights)
		fmt.Println(similarity)
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
