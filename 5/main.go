package main

import (
	"bufio"
	"flag"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := flag.String("i", "input", "input file name")
	flag.Parse()

	file, err := os.Open(*fileName)
	check(err)
	defer file.Close()

	m := make(map[int][]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), "|")
		if len(fields) < 2 {
			break
		}

		x, err := strconv.Atoi(fields[0])
		check(err)
		y, err := strconv.Atoi(fields[1])
		check(err)

		found := false
		for _, adj := range m[y] {
			if adj == x {
				found = true
				break
			}
		}
		if !found {
			m[y] = append(m[y], x)
		}
	}

	sum := 0
	sum2 := 0
	for scanner.Scan() {
		s := map[int]bool{}
		nums := castNums(strings.Split(scanner.Text(), ","))

		for _, num := range nums {
			s[num] = true
		}

		if isCorrect(nums, m, s) {
			sum += middleNumber(nums)
		} else {
			correct := correctlyOrder(nums, m, s)
			sum2 += middleNumber(correct)
		}
	}
	println(sum)
	println(sum2)
}

func correctlyOrder(nums []int, m map[int][]int, s map[int]bool) []int {
	ret := make([]int, 0, len(nums))
	for len(s) > 0 {
		keys := make([]int, 0, len(s))
		for k := range s {
			keys = append(keys, k)
		}

		for _, key := range keys {
			if !hasDeps(m, s, key) {
				ret = append(ret, key)
				delete(s, key)
			}
		}
	}
	return ret
}

func hasDeps(m map[int][]int, s map[int]bool, key int) bool {
	for _, next := range m[key] {
		if next != key && s[next] {
			return true
		}
	}
	return false
}

func middleNumber(nums []int) int {
	if len(nums) == 0 {
		panic("empty sequence")
	}

	return nums[len(nums)/2]
}

func contains(m map[int][]int, needle, curr int, s, seen map[int]bool) bool {
	if seen[curr] {
		return false
	}
	seen[curr] = true

	for _, next := range m[curr] {
		if needle == next {
			return true
		}

		if s[next] && contains(m, needle, next, s, seen) {
			return true
		}
	}

	return false
}

func isCorrect(nums []int, m map[int][]int, s map[int]bool) bool {
	for i := range len(nums) {
		for j := i + 1; j < len(nums); j++ {
			seen := map[int]bool{}
			if contains(m, nums[j], nums[i], s, seen) {
				return false
			}
		}
	}
	return true
}

func castNums(strs []string) []int {
	nums := make([]int, len(strs))
	for i := range strs {
		num, err := strconv.Atoi(strs[i])
		check(err)
		nums[i] = num
	}
	return nums
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
