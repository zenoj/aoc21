package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type line struct {
	x1, y1, x2, y2 int
}

func main() {
	// *. read in lines safe only if either x1 = x2 or y1 = y2
	// *. make int array[][] of size 1000*1000
	// *. enter the lines into the array
	// *. check array for position with value > 1
	// 1. read in boards
	data, err := os.ReadFile("/home/jay/go/src/challenges/aoc/ch5/puzzleInput.txt")
	check(err)
	inputLines := strings.Split(string(data), "\n")
	diagram := make([]int, 1000000)

	inputLines = strings.Split(string(data), "\n")
	for _, ILine := range inputLines {
		splitLine := strings.Split(ILine, " -> ")
		sVal := strings.Split(splitLine[0], ",")
		eVal := strings.Split(splitLine[1], ",")
		x1, _ := strconv.Atoi(sVal[0])
		x2, _ := strconv.Atoi(eVal[0])
		y1, _ := strconv.Atoi(sVal[1])
		y2, _ := strconv.Atoi(eVal[1])
		l := line{x1, y1, x2, y2}
		if x1 == x2 {
			placeLineInDiagram(diagram, "y", l)
		} else if y1 == y2 {
			placeLineInDiagram(diagram, "x", l)
		} else if isDiagonalTypeSum(l) {
			drawSumDiagonal(diagram, l)
		} else if isDiagonalTypeDiff(l) {
			drawDiffDiagonal(diagram, l)
		}
		continue
	}
	var c int
	for _, i2 := range diagram {
		if i2 > 1 {
			c++
		}
	}
	fmt.Printf("The number of danger spots is :%v", c)

}

func placeLineInDiagram(diagram []int, movingParameter string, l line) {
	maxX := int(math.Max(float64(l.x1), float64(l.x2)))
	minX := int(math.Min(float64(l.x1), float64(l.x2)))
	maxY := int(math.Max(float64(l.y1), float64(l.y2)))
	minY := int(math.Min(float64(l.y1), float64(l.y2)))
	if movingParameter == "y" {
		for i := minY; i <= maxY; i++ {
			diagram[transIndex(l.x1, i, 1000)]++
		}
	} else {
		for i := minX; i <= maxX; i++ {
			diagram[transIndex(i, l.y1, 1000)]++
		}
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func transIndex(x, y, rowSize int) int {
	return x + y*rowSize
}

func isDiagonalTypeSum(l line) bool {
	// a line is diagonal type "sum" if the sum of both coordinates is the same.
	return l.x1+l.y1 == l.x2+l.y2
}

func isDiagonalTypeDiff(l line) bool {
	return math.Abs(float64(l.x1)-float64(l.y1)) == math.Abs(float64(l.x2)-float64(l.y2))
}

func drawSumDiagonal(diagram []int, l line) {
	minX := int(math.Min(float64(l.x1), float64(l.x2)))
	maxX := int(math.Max(float64(l.x1), float64(l.x2)))
	maxY := int(math.Max(float64(l.y1), float64(l.y2)))
	for startX := minX; startX <= maxX; startX++ {
		diagram[transIndex(startX, maxY, 1000)]++
		maxY--
	}
}

func drawDiffDiagonal(diagram []int, l line) {
	minX := int(math.Min(float64(l.x1), float64(l.x2)))
	maxX := int(math.Max(float64(l.x1), float64(l.x2)))
	minY := int(math.Min(float64(l.y1), float64(l.y2)))

	for startX := minX; startX <= maxX; startX++ {
		diagram[transIndex(startX, minY, 1000)]++
		minY++
	}
}
