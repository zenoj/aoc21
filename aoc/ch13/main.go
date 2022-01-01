package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var maxY int = 895

//var maxY int = 15

var maxX int = 1311

//var maxX int = 11

func main() {

	// read in the puzzle input
	data, err := os.ReadFile("/home/jay/go/src/challenges/aoc/ch13/puzzleInput.txt")
	check(err)
	inputLines := strings.Split(string(data), "\n\n")
	dotPositions := strings.Split(inputLines[0], "\n")
	foldInstructions := strings.Split(inputLines[1], "\n")
	m := make([]bool, maxX*maxY)

	// draw dots into array
	for _, position := range dotPositions {
		pos := strings.Split(position, ",")
		x, _ := strconv.Atoi(pos[0])
		y, _ := strconv.Atoi(pos[1])
		m[transIndex(x, y, maxX)] = true
	}

	// folding: take out first fold instruction
	for i := 0; i < 12; i++ {
		fLine := strings.Split(foldInstructions[i], "fold along ")[1]
		if string(fLine[0]) == "y" {
			horizontalLine, _ := strconv.Atoi(fLine[2:])
			m = foldOnY(m, horizontalLine)
		} else {
			verticalLine, _ := strconv.Atoi(fLine[2:])
			m = foldOnX(m, verticalLine)
		}
	}

	// draw the code
	for i, element := range m {
		if i%maxX == 0 {
			fmt.Printf("\n")
		}
		if element {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}
	}
}

func foldOnY(m []bool, yValue int) []bool {
	for y := yValue + 1; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if m[transIndex(x, y, maxX)] {
				// mirror the dot
				newY := int(float64(yValue) - (math.Abs(float64(y) - float64(yValue))))

				m[transIndex(x, newY, maxX)] = true
			}
		}
	}
	// throw away the rows bigger than yValue
	m = m[:transIndex(0, yValue, maxX)]
	maxY = yValue
	return m
}

func foldOnX(m []bool, xValue int) []bool {
	for y := 0; y < maxY; y++ {
		for x := xValue + 1; x < maxX; x++ {
			if m[transIndex(x, y, maxX)] {
				// mirror the dot
				newX := int(float64(xValue) - (math.Abs(float64(x) - float64(xValue))))
				m[transIndex(newX, y, maxX)] = true
			}
		}
	}
	// throw away the rows bigger than yValue
	m = cutOffAt(m, xValue)
	maxX = xValue
	return m
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func transIndex(x, y, rowSize int) int {
	return x + y*rowSize
}

func cutOffAt(m []bool, xValue int) []bool {
	shortenedArray := make([]bool, 0, xValue*maxY)
	for i := 0; i < maxY; i++ {
		startIndex := transIndex(0, i, maxX)
		endIndex := transIndex(xValue, i, maxX)
		shortenedArray = append(shortenedArray, m[startIndex:endIndex]...)
	}
	return shortenedArray
}
