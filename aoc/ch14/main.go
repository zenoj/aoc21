package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {

	// read in the puzzle input
	data, err := os.ReadFile("/home/jay/go/src/challenges/aoc/ch14/puzzleInput.txt")
	check(err)
	inputLines := strings.Split(string(data), "\n\n")
	template := inputLines[0]
	rules := strings.Split(inputLines[1], "\n")
	ruleMap := make(map[string]string)

	newFrequencyMap := make(map[string]int)
	oldFrequencyMap := make(map[string]int)
	elementFrequencyMap := make(map[string]int)

	// insert rules into map
	for _, position := range rules {
		rule := strings.Split(position, " -> ")
		ruleMap[rule[0]] = rule[1]
	}

	// initialize newFrequencyMap with template
	for pair := 0; pair < len(template)-1; pair += 1 {
		oldFrequencyMap[template[pair:pair+2]]++
	}

	// populate
	numSteps := 40
	for steps := 0; steps < numSteps; steps++ {
		oldFrequencyMap, newFrequencyMap = proliferate(oldFrequencyMap, newFrequencyMap, ruleMap)
	}

	// count elements in map
	for rule, freq := range oldFrequencyMap {
		elementFrequencyMap[string(rule[0])] += freq
		elementFrequencyMap[string(rule[1])] += freq
	}

	// find max element in map
	maxValue := 0
	maxKey := ""
	minValue := math.MaxInt
	minKey := ""
	for k, v := range elementFrequencyMap {
		if v > maxValue {
			maxValue = v
			maxKey = k
		} else if v < minValue {
			minValue = v
			minKey = k
		}
	}
	if maxKey == string(template[0]) {
		maxValue++
	}

	if maxKey == string(template[len(template)-1]) {
		maxValue++
	}

	if minKey == string(template[0]) {
		minValue++
	}

	if minKey == string(template[len(template)-1]) {
		minValue++
	}
	minValue = int(math.Round(float64(minValue) / 2))
	maxValue = int(math.Round(float64(maxValue) / 2))

	fmt.Printf("Difference : %v\n", maxValue-minValue)
	fmt.Printf("MinValue: %v\n", minValue)
	fmt.Printf("MaxValue %v", maxValue)

}

func proliferate(oldMap, newMap map[string]int, ruleMap map[string]string) (map[string]int, map[string]int) {
	for k, v := range oldMap {
		newMap[string(k[0])+ruleMap[k]] += v
		newMap[ruleMap[k]+string(k[1])] += v
	}
	oldMap = make(map[string]int)
	return newMap, oldMap
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
