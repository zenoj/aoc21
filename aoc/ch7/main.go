package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// *. sort array
	// *. find median
	// *. if no median exists take average of the two numbers in the middle
	// *. sum up the differences of every number to the chosen middle position
	data, err := os.ReadFile("/home/jay/go/src/challenges/aoc/ch7/puzzleInput.txt")
	check(err)
	inputStrings := strings.Split(string(data), ",")
	positions := make([]int, 0)
	for _, p := range inputStrings {
		i, _ := strconv.Atoi(p)
		positions = append(positions, i)
	}

	sort.Ints(positions)

	var c int
	var currentChosen int
	var min int = math.MaxInt
	for chosen := 0; chosen <= positions[len(positions)-1]; chosen++ {
		for _, position := range positions {
			c += int(sum(1, int(math.Abs(float64(position)-float64(chosen)))))
		}
		if c < min {
			min = c
			currentChosen = chosen
		}
		c = 0
	}

	fmt.Printf("The number of fuel needed is: %v with the chosen number being: %v", min, currentChosen)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sum(s, e int) int {
	var c int
	for i := s; i <= e; i++ {
		c += i
	}
	return c
}
