package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("/home/jay/go/src/challenges/aoc/ch6/puzzleInput.txt")
	check(err)
	fishStrings := strings.Split(string(data), ",")
	oldPopulation := make(map[int]int)
	newPopulation := make(map[int]int)
	// put in the fish with their day values
	for _, f := range fishStrings {
		fish, _ := strconv.Atoi(f)
		oldPopulation[fish]++
	}

	var numberOfDays int = 256
	for i := 0; i < numberOfDays; i++ {
		oldPopulation, newPopulation = aDayGoesBy(newPopulation, oldPopulation)
	}
	var c int
	for _, value := range newPopulation {
		c += value
	}
	fmt.Printf("The number of fish is: %v", c)
}

func aDayGoesBy(newPopulation, oldPopulation map[int]int) (map[int]int, map[int]int) {
	for i := 0; i <= 7; i++ {
		newPopulation[i] = oldPopulation[i+1]
	}
	newPopulation[8] = oldPopulation[0]
	newPopulation[6] += oldPopulation[0]
	copyMap(newPopulation, oldPopulation)
	return oldPopulation, newPopulation
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func copyMap(src, dest map[int]int) {
	for key, value := range src {
		dest[key] = value
	}
}
