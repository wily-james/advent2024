package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

func main() {
	input := flag.String("i", "input", "input file name")
	partTwo := flag.Bool("2", false, "run part two")
	flag.Parse()

	bytes, err := os.ReadFile(*input)
	check(err)

	fs := make([]int, 0, len(bytes))
	for i := range bytes {
		if bytes[i] == ' ' || bytes[i] == '\n' || bytes[i] == 't' || bytes[i] == '\r' {
			break
		}
		if bytes[i] < '0' || bytes[i] > '9' {
			panic(fmt.Errorf("Got invalid digit %d", bytes[i]))
		}
		fs = append(fs, int(bytes[i])-'0')
	}

	if *partTwo {
		fmt.Println(compact2(fs))
	} else {
		fmt.Println(compact(fs))
	}
}

type Space struct {
	loc   int
	space int
}

func compact2(fs []int) int {
	if len(fs) <= 1 {
		return 0
	}

	fileLocs := make([]int, len(fs)/2+1)
	fileSizes := make([]int, len(fs)/2+1)
	spaceLocs := make([]Space, len(fs)/2+1)
	blockCount := fs[0]
	for fi := 2; fi < len(fs); fi += 2 {
		spaceLocs[(fi-1)/2] = Space{
			blockCount,
			fs[fi-1],
		}
		blockCount += fs[fi-1]
		fileLocs[fi/2] = blockCount
		fileSizes[fi/2] = fs[fi]
		blockCount += fs[fi]
	}
	fmt.Printf("len(fs): %d\n", len(fs))
	fmt.Printf("FileLocs: %+v\n", fileLocs)
	fmt.Printf("FileSizes: %+v\n", fileSizes)
	fmt.Printf("SpaceLocs: %+v\n", spaceLocs)

	start := len(fs) - 1
	if len(fs)%2 == 0 {
		start = len(fs) - 2
	}
	for fi := start; fi >= 0; fi -= 2 {
		for si := range spaceLocs {
			if spaceLocs[si].loc >= fileLocs[fi/2] {
				continue
			}
			if spaceLocs[si].space < fs[fi] {
				continue
			}

			spaceLocs = append(spaceLocs, Space{
				fileLocs[fi/2],
				fs[fi],
			})
			fileLocs[fi/2] = spaceLocs[si].loc
			spaceLocs[si].loc += fs[fi]
			spaceLocs[si].space -= fs[fi]
			spaceLocs = compactSpace(spaceLocs)
			break
		}
	}

	checkSum := 0
	fmt.Printf("FileLocs: %+v\n", fileLocs)
	fmt.Printf("FileSizes: %+v\n", fileSizes)
	fmt.Printf("SpaceLocs: %+v\n", spaceLocs)

	for id, loc := range fileLocs {
		blockCount := loc
		for range fs[id*2] {
			checkSum += id * blockCount
			blockCount++
		}
	}

	return checkSum
}

func compactSpace(spaces []Space) []Space {
	sort.Slice(spaces, func(i, j int) bool {
		return spaces[i].loc < spaces[j].loc
	})
	for i := len(spaces) - 1; i >= 1; i-- {
		if spaces[i-1].loc+spaces[i-1].space >= spaces[i].loc {
			spaces[i-1].space = (spaces[i].loc + spaces[i].space) - spaces[i-1].loc
			spaces[i].space = 0
		}
	}

	newSpaces := make([]Space, 0, len(spaces))
	for _, space := range spaces {
		if space.space > 0 {
			newSpaces = append(newSpaces, space)
		}
	}
	return newSpaces
}

func compact(fs []int) int {
	if len(fs) <= 1 {
		return 0
	}

	checkSum := 0
	si, fi := 1, len(fs)-1
	space, blocks := fs[si], fs[fi]
	blockCount := fs[0]
	for si < fi {
		for blocks > 0 {
			spaceToUse := min(space, blocks)
			for range spaceToUse {
				checkSum += fi / 2 * blockCount
				blockCount++
			}
			blocks -= spaceToUse
			space -= spaceToUse

			if space == 0 {
				if si+1 == fi {
					for range blocks {
						checkSum += fi / 2 * blockCount
						blockCount++
					}
					return checkSum
				} else {
					for range fs[si+1] {
						checkSum += (si + 1) / 2 * blockCount
						blockCount++
					}
				}
				si += 2
				space = fs[si]
			}
		}

		if si == fi-1 {
			return checkSum
		}

		fi -= 2
		blocks = fs[fi]
	}

	return checkSum
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
