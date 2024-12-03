package main

import (
	"flag"
	"os"
)

type State int

const (
	Start State = iota
	M
	U
	L
	LPAREN
	N1

	D
	O
	DOLPAREN
	N
	APOS
	T
	DONTLPAREN
)

type Status struct {
	state     State
	currDig   int
	digCount  int
	firstDig  int
	secondDig int
	enabled   bool
}

func main() {
	fileName := flag.String("i", "input", "input file name")
	doPart2 := flag.Bool("2", false, "pass to run part 2")
	flag.Parse()

	buf, err := os.ReadFile(*fileName)
	check(err)

	curr := 0
	state := Status{}
	state.enabled = true
	for i := 0; i < len(buf); i++ {
		b := buf[i]
		switch state.state {
		case Start:
			if b == 'm' {
				state.state = M
				continue
			} else if *doPart2 && b == 'd' {
				state.state = D
				continue
			}
			i++
		case D:
			if b == 'o' {
				state.state = O
				continue
			}
		case O:
			if b == '(' {
				state.state = DOLPAREN
				continue
			}
			if b == 'n' {
				state.state = N
				continue
			}
		case DOLPAREN:
			if b == ')' {
				state.enabled = true
				i++
			}
		case N:
			if b == '\'' {
				state.state = APOS
				continue
			}
		case APOS:
			if b == 't' {
				state.state = T
				continue
			}
		case T:
			if b == '(' {
				state.state = DONTLPAREN
				continue
			}
		case DONTLPAREN:
			if b == ')' {
				state.enabled = false
				i++
			}
		case M:
			if b == 'u' {
				state.state = U
				continue
			}
		case U:
			if b == 'l' {
				state.state = L
				continue
			}
		case L:
			if b == '(' {
				state.state = LPAREN
				continue
			}
		case LPAREN:
			if isDigit(b) {
				state.currDig *= 10
				state.currDig += int(b) - '0'
				state.digCount += 1
				if state.digCount > 3 {
					break
				}
				continue
			} else if b == ',' {
				if state.digCount < 1 {
					break
				}
				state.firstDig = state.currDig
				state.currDig = 0
				state.digCount = 0
				state.state = N1
				continue
			}
		case N1:
			if isDigit(b) {
				state.currDig *= 10
				state.currDig += int(b) - '0'
				state.digCount += 1
				if state.digCount > 3 {
					break
				}
				continue
			} else if b == ')' {
				if state.digCount < 1 {
					break
				}
				if state.enabled {
					curr += state.firstDig * state.currDig
				}
				i++
			}
		}
		enabled := state.enabled
		state = Status{}
		state.enabled = enabled
		i--
	}
	println(curr)
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
