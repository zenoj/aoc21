package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var counter int
	f, err := os.Open("/home/jay/go/src/challenges/aoc/ch1/puzzleInput.txt")
	check(err)
	s := bufio.NewScanner(f)
	s.Scan()
	aa, _ := strconv.Atoi(s.Text())
	s.Scan()
	bb, _ := strconv.Atoi(s.Text())
	s.Scan()
	cc, _ := strconv.Atoi(s.Text())
	oldValue := aa + bb + cc
	for {
		n := s.Scan()
		if !n {
			break
		}
		nn, _ := strconv.Atoi(s.Text())
		newValue := bb + cc + nn
		if newValue > oldValue {
			counter++
		}
		bb = cc
		cc = nn
		oldValue = newValue
	}
	fmt.Printf("The number of smaller values is: %d", counter)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
