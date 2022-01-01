package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	data, err := os.ReadFile("/home/jay/go/src/challenges/aoc/ch3/puzzleInput.txt")
	check(err)
	stringData := string(data)
	oxygonGen := strings.Split(stringData, "\n")
	restOxygen := oxygonGen
	co2scrubber := strings.Split(stringData, "\n")

	inputLength := len(oxygonGen[0])
	for i := 0; i < inputLength; i++ {
		commonBit := computeMostCommonBit(restOxygen, i)
		restOxygen = filterOxygen(restOxygen, i, commonBit)
		if len(restOxygen) == 1 {
			break
		}
	}
	fmt.Printf("The Oxygon Generator rating is : %v", restOxygen)

	for i := 0; i < inputLength; i++ {
		commonBit := computeMostCommonBit(co2scrubber, i)
		co2scrubber = filterCO2(co2scrubber, i, commonBit)
		if len(co2scrubber) == 1 {
			break
		}
	}
	fmt.Printf("The CO2 scrubber rating is : %v", co2scrubber)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func computeMostCommonBit(stringData []string, pos int) int {
	cCommon := 0
	var c int
	for _, i := range stringData {
		t := string(i[pos])
		tmp, _ := strconv.Atoi(t)
		cCommon += tmp
		c++
	}

	if cCommon > c/2 {
		return 1
	}
	if cCommon == c/2 && c%2 == 0 {
		return 2
	}
	return 0
}

func filterOxygen(stringData []string, pos int, mostCommon int) []string {
	if mostCommon == 2 {
		mostCommon = 1
	}
	mC := mostCommon
	for i := 0; i < len(stringData); i++ {
		tmpStr := stringData[i]
		e, _ := strconv.Atoi(string(tmpStr[pos]))
		if e != mC {
			stringData = remove(stringData, i)
			i--
		}
	}
	return stringData
}

func filterCO2(stringData []string, pos int, mostCommon int) []string {

	if mostCommon == 0 {
		mostCommon = 1
	} else if mostCommon == 1 {
		mostCommon = 0
	} else if mostCommon == 2 {
		mostCommon = 0
	}
	mC := mostCommon
	for i := 0; i < len(stringData); i++ {
		tmpStr := stringData[i]
		e, _ := strconv.Atoi(string(tmpStr[pos]))
		if e != mC {
			stringData = remove(stringData, i)
			i--
		}
	}
	return stringData
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
