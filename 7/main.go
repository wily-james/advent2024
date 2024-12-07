package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Equation struct {
	result   uint64
	operands []uint64
}

func main() {
	input := flag.String("i", "input", "input file")
	partTwo := flag.Bool("2", false, "pass for part two")
	flag.Parse()

	file, err := os.Open(*input)
	check(err)
	defer file.Close()

	equations := make([]Equation, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		eq := Equation{}
		i := 0
		for ; i < len(bytes) && isDigit(bytes[i]); i++ {
			eq.result *= 10
			eq.result += digToInt(bytes[i])
		}

		i++
		for i < len(bytes) {
			num := uint64(0)
			for i++; i < len(bytes) && isDigit(bytes[i]); i++ {
				num *= 10
				num += digToInt(bytes[i])
			}
			eq.operands = append(eq.operands, num)
		}
		equations = append(equations, eq)
	}

	sum := uint64(0)
	for _, eq := range equations {
		if isValid(eq, *partTwo) {
			lastSum := sum
			sum += eq.result
			if sum < lastSum {
				panic("Overflow!")
			}
		}
	}
	fmt.Println(sum)
}

func isValid(eq Equation, useConcat bool) bool {
	if len(eq.operands) == 0 {
		return false
	}
	return isValidHelper(eq, 1, eq.operands[0], useConcat)
}

func isValidHelper(eq Equation, i int, total uint64, useConcat bool) bool {
	if i >= len(eq.operands) {
		return total == eq.result
	}

	if isValidHelper(eq, i+1, total+eq.operands[i], useConcat) {
		return true
	}

	if isValidHelper(eq, i+1, total*eq.operands[i], useConcat) {
		return true
	}

	if !useConcat {
		return false
	}

	res, err := concat(total, eq.operands[i])
	if err != nil {
		return false
	}
	return isValidHelper(eq, i+1, res, useConcat)
}

func concat(l, r uint64) (uint64, error) {
	sl, sr := l, r

	var flipped uint64 = 1
	for r > 0 {
		flipped *= 10
		flipped += r % 10
		r /= 10
	}

	for flipped > 1 {
		l *= 10
		l += flipped % 10
		flipped /= 10
	}

	if l < sl || l < sr {
		return 0, fmt.Errorf("Overflowed!")
	}
	return l, nil
}

func digToInt(b byte) uint64 {
	return uint64(b) - '0'
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
