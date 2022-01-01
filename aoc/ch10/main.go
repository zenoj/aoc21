package main

import (
	"aoc.com/utils/dataStructures/stringStack"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// read in data
	data, err := os.ReadFile("/home/jay/go/src/challenges/aoc/ch10/puzzleInput.txt")
	check(err)
	inputLines := strings.Split(string(data), "\n")
	numLines := len(inputLines)
	m := make([][]string, numLines)
	for i, line := range inputLines {
		m[i] = strings.Split(line, "")
	}
	bracketsMap := make(map[string]string)
	bracketsMap[")"] = "("
	bracketsMap["]"] = "["
	bracketsMap[">"] = "<"
	bracketsMap["}"] = "{"
	bracketsMap["("] = ")"
	bracketsMap["["] = "]"
	bracketsMap["<"] = ">"
	bracketsMap["{"] = "}"

	scoreMap := make(map[string]int)
	scoreMap[")"] = 1
	scoreMap["]"] = 2
	scoreMap["}"] = 3
	scoreMap[">"] = 4

	var syntaxErrorScores []int
	for _, characterList := range m {
		if isCorrupted(characterList, bracketsMap) {
			continue
		}
		chars := completeLines(characterList, bracketsMap)
		syntaxErrorScores = append(syntaxErrorScores, computeScore(chars))
	}
	sort.Ints(syntaxErrorScores)
	fmt.Printf("%v", syntaxErrorScores[len(syntaxErrorScores)/2])
}

func completeLines(characterList []string, bracketsMap map[string]string) []string {
	s := stringStack.Stack{}
	reverseStringList := make([]string, 0)
	for _, c := range characterList {
		if isOpeningBraces(c) {
			s.Push(c)
		} else if bracketsMap[c] == s.Top() {
			s.Pop()
			continue
		}
	}
	size := s.Size()
	for e := 0; e < size; e++ {
		reverseStringList = append(reverseStringList, bracketsMap[s.Pop()])
	}
	return reverseStringList
}

func isCorrupted(characterList []string, bracketsMap map[string]string) bool {
	s := stringStack.Stack{}
	for _, c := range characterList {
		if isOpeningBraces(c) {
			s.Push(c)
		} else if bracketsMap[c] == s.Top() {
			s.Pop()
			continue
		} else {
			return true
		}

	}
	return false
}

func isOpeningBraces(c string) bool {
	switch c {
	case "[", "{", "<", "(":
		return true
	}
	return false
}

func computeScore(characterList []string) int {

	scoreMap := make(map[string]int)
	scoreMap[")"] = 1
	scoreMap["]"] = 2
	scoreMap["}"] = 3
	scoreMap[">"] = 4

	score := 0

	for _, s := range characterList {
		score *= 5
		score += scoreMap[s]
	}
	return score
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
