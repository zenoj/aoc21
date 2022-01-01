package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("/home/jay/go/src/challenges/aoc/ch2/puzzleInput.txt")
	check(err)
	s := bufio.NewScanner(f)
	var depth int
	var horizontal int
	var aim int
	for s.Scan() {
		a := strings.Split(s.Text(), " ")[0]
		b, _ := strconv.Atoi(strings.Split(s.Text(), " ")[1])
		switch a {
		case "forward":
			horizontal += b
			depth += aim * b

		case "down":
			aim += b
		case "up":
			aim -= b
		}
	}
	fmt.Printf("The value is: %d", depth*horizontal)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
