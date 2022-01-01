package main

import (
	"aoc.com/utils/mathHelp"
	"aoc.com/utils/str"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type pos struct {
	row, column int
}

func main() {
	// read in data
	data, err := os.ReadFile("/home/jay/go/src/challenges/aoc/ch9/puzzleInput.txt")
	check(err)
	inputLines := strings.Split(string(data), "\n")
	numLines := len(inputLines)
	numColumns := len(inputLines[0])
	m := make([][]int, numLines)
	for i, line := range inputLines {
		m[i] = str.StrArrayToIntArray(strings.Split(line, ""))
	}
	// find minimums
	mins := make([]pos, 0)
	for r := 0; r < numLines; r++ {
		for c := 0; c < numColumns; c++ {
			if isLowPoint(m, r, c) {
				mins = append(mins, pos{r, c})
			}
		}
	}
	fmt.Printf("%v", mins)
	basins := make([][]pos, len(mins))
	for i, min := range mins {
		visited := make(map[string]pos, 0)
		visited[min.toString()] = min
		basins[i] = append(buildBasins(min, m, visited))
		basins[i] = append(basins[i], min)

	}
	// remove duplicates from slices

	sizeSlice := make([]int, len(basins))
	// find biggest basins
	// make list with basins size
	// sort basins size list
	for i, basin := range basins {
		sizeSlice[i] = len(basin)
	}
	fmt.Printf("%v", basins)
	sort.Ints(sizeSlice)
	threeBiggest := sizeSlice[len(sizeSlice)-3:]
	fmt.Printf("%v", threeBiggest)
	fmt.Printf("%v", mathHelp.Product(threeBiggest))
}

func isLowPoint(m [][]int, r, c int) bool {
	point := m[r][c]
	if c > 0 {
		if m[r][c-1] <= point {
			return false
		}
	}
	// get right neighbor
	if c < len(m[0])-1 {
		if m[r][c+1] <= point {
			return false
		}
	}
	// get upper neighbor
	if r > 0 {
		if m[r-1][c] <= point {
			return false
		}
	}
	// get lower neighbor
	if r < len(m)-1 {
		if m[r+1][c] <= point {
			return false
		}
	}
	return true
}

func getNeighborsNot9(m [][]int, r, c int, visited map[string]pos) []pos {
	neighbors := make([]pos, 0)
	// get left neighbor

	if c > 0 && m[r][c-1] != 9 {
		neighbors = append(neighbors, pos{r, c - 1})
	}
	// get right neighbor
	if c < len(m[0])-1 && m[r][c+1] != 9 {
		neighbors = append(neighbors, pos{r, c + 1})
	}
	// get upper neighbor
	if r > 0 && m[r-1][c] != 9 {
		neighbors = append(neighbors, pos{r - 1, c})
	}
	// get lower neighbor
	if r < len(m)-1 && m[r+1][c] != 9 {
		neighbors = append(neighbors, pos{r + 1, c})
	}
	biggerNeighbors := make([]pos, 0)
	for _, neighbor := range neighbors {
		key := neighbor.toString()
		if _, ok := visited[key]; !ok {
			visited[key] = neighbor
			biggerNeighbors = append(biggerNeighbors, neighbor)
		}
	}
	return biggerNeighbors
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func buildBasins(lowPoint pos, m [][]int, visited map[string]pos) []pos {
	neighbors := getNeighborsNot9(m, lowPoint.row, lowPoint.column, visited)
	for _, neighbor := range neighbors {
		neighbors = append(neighbors, buildBasins(neighbor, m, visited)...)
	}
	return neighbors
}

func (p *pos) toString() string {
	r := strconv.Itoa(p.row)
	c := strconv.Itoa(p.column)
	return r + " " + c
}
