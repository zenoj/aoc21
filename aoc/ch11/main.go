package main

import (
	"aoc.com/utils/str"
	"fmt"
	"os"
	"strings"
)

type pos struct {
	row, column int
}

func main() {
	data, err := os.ReadFile("/home/jay/go/src/challenges/aoc/ch11/puzzleInput.txt")
	check(err)
	inputLines := strings.Split(string(data), "\n")
	numLines := len(inputLines)
	m := make([][]int, numLines)
	for i, line := range inputLines {
		m[i] = str.StrArrayToIntArray(strings.Split(line, ""))
	}
	var numFlashes int
	s := 0
	// for the first part just add all the flashes up and print that
	for ; ; s++ {
		numFlashes = step(m)
		if numFlashes == len(m)*len(m[0]) {
			break
		}
	}
	fmt.Printf("%v", s+1)
}

func step(m [][]int) int {
	var counter int
	hasFlashedThisStep := make([][]bool, len(m))
	for i := 0; i < len(m); i++ {
		hasFlashedThisStep[i] = make([]bool, len(m[0]))
	}
	// increase energy level of each octopus
	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m[0]); c++ {
			m[r][c]++
		}
	}
	// flash until everybody who could, did flash.
	someOneFlashed := true
	for someOneFlashed {
		someOneFlashed = false
		for r := 0; r < len(m); r++ {
			for c := 0; c < len(m[0]); c++ {
				if m[r][c] > 9 && !hasFlashedThisStep[r][c] {
					flash(m, r, c)
					hasFlashedThisStep[r][c] = true
					counter++
					someOneFlashed = true
				}
			}
		}
	}
	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m[0]); c++ {
			if hasFlashedThisStep[r][c] {
				m[r][c] = 0
			}
		}
	}
	return counter
}

func flash(m [][]int, r, c int) {
	// 1. get all neighbors
	neighbors := getNeighbors(m, r, c)
	// increase all neighbors energy levels by 1
	for _, n := range neighbors {
		m[n.row][n.column]++
	}
}

func getNeighbors(m [][]int, r, c int) []pos {
	neighbors := make([]pos, 0)
	// get left neighbor and SW
	if c > 0 {
		neighbors = append(neighbors, pos{r, c - 1})
		if r < len(m)-1 {
			neighbors = append(neighbors, pos{r + 1, c - 1})
		}
	}
	// get right neighbor and NE
	if c < len(m[0])-1 {
		neighbors = append(neighbors, pos{r, c + 1})
		if r > 0 {
			neighbors = append(neighbors, pos{r - 1, c + 1})
		}
	}
	// get upper neighbor and NW
	if r > 0 {
		neighbors = append(neighbors, pos{r - 1, c})
		if c > 0 {
			neighbors = append(neighbors, pos{r - 1, c - 1})
		}
	}
	// get lower neighbor and SE
	if r < len(m)-1 {
		neighbors = append(neighbors, pos{r + 1, c})
		if c < len(m[0])-1 {
			neighbors = append(neighbors, pos{r + 1, c + 1})
		}
	}
	return neighbors
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
