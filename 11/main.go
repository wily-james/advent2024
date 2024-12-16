package main

import (
	"flag"
	"fmt"
	"os"
)

type Key struct {
	stone  int
	blinks int
}

func numStones(stone, blinks int, memo map[Key]int) int {
	if blinks == 0 {
		return 1
	}

	key := Key{stone, blinks}
	res, ok := memo[key]
	if ok {
		return res
	}

	switch {
	case stone == 0:
		res = numStones(1, blinks-1, memo)
	case numDigits(stone)%2 == 0:
		left, right := split(stone)
		res = numStones(left, blinks-1, memo) + numStones(right, blinks-1, memo)
	default:
		res = numStones(stone*2024, blinks-1, memo)
	}

	memo[key] = res
	return res
}

func main() {
	input := flag.String("i", "input", "input file")
	blinks := flag.Int("b", 25, "number of blinks")
	flag.Parse()

	bytes, err := os.ReadFile(*input)
	check(err)

	//fmt.Printf("b: %d\n", *blinks)
	stones := parseStones(bytes)
	//fmt.Printf("%+v\n", stones)
	sum := 0
	memo := make(map[Key]int)
	for _, stone := range stones {
		sum += numStones(stone, *blinks, memo)
	}

	//fmt.Printf("%+v\n", stones)
	//fmt.Printf("%d\n", len(stones))
	fmt.Printf("%d\n", sum)
}

func split(n int) (int, int) {
	rev := 1
	numDigits := 0
	for n != 0 {
		rev *= 10
		rev += n % 10
		n /= 10
		numDigits++
	}

	left := 0
	for range numDigits / 2 {
		left *= 10
		left += rev % 10
		rev /= 10
	}

	right := 0
	for range numDigits / 2 {
		right *= 10
		right += rev % 10
		rev /= 10
	}

	return left, right
}

func numDigits(n int) int {
	if n == 0 {
		return 1
	}

	count := 0
	for n != 0 {
		n /= 10
		count++
	}

	return count
}

func parseStones(bytes []byte) []int {
	stones := make([]int, 0)
	num := -1
	for _, byte := range bytes {
		if byte >= '0' && byte <= '9' {
			dig := int(byte) - '0'
			if num < 0 {
				num = dig
			} else {
				num *= 10
				num += dig
			}
		} else if num >= 0 {
			stones = append(stones, num)
			num = -1
		}
	}
	if num >= 0 {
		stones = append(stones, num)
	}
	return stones
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
