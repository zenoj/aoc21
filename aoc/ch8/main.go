package main

import (
	"aoc.com/utils/str"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("/home/jay/go/src/challenges/aoc/ch8/puzzleInput.txt")
	check(err)
	inputLines := strings.Split(string(data), "\n")
	var c int
	for _, line := range inputLines {
		outputLine := strings.Split(line, " | ")
		c += extract(outputLine[0], outputLine[1])
	}
	fmt.Println(c)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func extract(sigPatLine string, outputLine string) int {
	patternMapping := make(map[int]string)
	letterMapping := make(map[string]string)
	sigPatList := strings.Fields(sigPatLine)
	fiveLetters := make([]string, 0)
	sixLetters := make([]string, 0)
	// initialize mappings with 1, 4, 7, 8
	for _, p := range sigPatList {
		if len(p) == 2 {
			patternMapping[1] = p
		} else if len(p) == 3 {
			patternMapping[7] = p
		} else if len(p) == 4 {
			patternMapping[4] = p
		} else if len(p) == 7 {
			patternMapping[8] = p
		} else if len(p) == 5 {
			fiveLetters = append(fiveLetters, p)
		} else if len(p) == 6 {
			sixLetters = append(sixLetters, p)
		}
	}
	letterMapping["a"] = str.Diff(patternMapping[7], patternMapping[1])
	// compute intersection of all with length 5
	inter5_12 := str.Intersection(fiveLetters[0], fiveLetters[1])
	inter5_123 := str.Intersection(inter5_12, fiveLetters[2])
	dg := str.Diff(inter5_123, letterMapping["a"])
	letterMapping["d"] = str.Intersection(dg, patternMapping[4])
	letterMapping["g"] = str.Diff(dg, letterMapping["d"])
	letterMapping["b"] = str.Diff(patternMapping[4], patternMapping[1], letterMapping["d"])
	letterMapping["e"] = str.Diff(patternMapping[8], letterMapping["a"], letterMapping["g"], patternMapping[4])
	// now we find 6ers
	for _, sigPat := range sixLetters {
		if str.Diff(patternMapping[8], sigPat) == letterMapping["d"] {
			patternMapping[0] = sigPat
		} else if str.Diff(patternMapping[8], sigPat) == letterMapping["e"] {
			patternMapping[9] = sigPat
		} else {
			patternMapping[6] = sigPat
			letterMapping["c"] = str.Diff(patternMapping[8], sigPat)
		}
	}
	letterMapping["f"] = str.Diff(patternMapping[1], letterMapping["c"])
	// lastly we find 5ers
	for _, sigPat := range fiveLetters {
		if strings.Contains(sigPat, letterMapping["b"]) {
			patternMapping[5] = sigPat
		} else if strings.Contains(sigPat, letterMapping["e"]) {
			patternMapping[2] = sigPat
		} else {
			patternMapping[3] = sigPat
		}
	}
	// now we have both mappings complete we can start deciphering the output side
	// for that we just build the intersection and see if the return value is of the same length
	// as the current number we are querying for
	outputValues := strings.Fields(outputLine)
	decipheredValues := make([]string, 0)
	for _, outputValue := range outputValues {
		for k, v := range patternMapping {
			if len(v) == len(outputValue) {
				if len(str.Intersection(v, outputValue)) == len(outputValue) {
					decipheredValues = append(decipheredValues, strconv.Itoa(k))
					break
				}
			}
		}
	}
	a, _ := strconv.Atoi(strings.Join(decipheredValues, ""))
	return a

}
