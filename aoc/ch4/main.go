package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type table struct {
	numbersTable [5][]string
	successTable [5][]bool
}
type position struct {
	row    string
	column string
}

func main() {
	// * read in boards
	// * save 2 instances of each board, one with the actual numbers and another one to fill in finished cells
	// * simulate game and have a won() function that works on boards
	// * always safe the last proposed number

	// 1. read in boards
	data, err := os.ReadFile("/home/jay/go/src/challenges/aoc/ch4/demoInput.txt")
	check(err)
	inputData := strings.Split(string(data), "\n\n")
	numbers := inputData[0]
	nField := strings.Split(numbers, ",")
	tableData := inputData[1:]
	var tables = make([]table, len(tableData))
	var lines []string
	// fill tables
	for i, table := range tableData {
		lines = strings.Split(table, "\n")
		for i2, line := range lines {
			tables[i].numbersTable[i2] = strings.Fields(line)
			tables[i].successTable[i2] = make([]bool, 5)
		}
	}

	// * simulate game
	t, lastNum := run(tables, nField)
	score := t.computeScore(lastNum)
	fmt.Println(score)
}

func run(tables []table, numbers []string) (table, string) {
	// runs the bingo game till the end
	var currentNumber string
	var contains bool
	var won bool
	var pos position
	for i := 0; i < len(numbers); i++ {
		currentNumber = numbers[i]
		for j := 0; j < len(tables); j++ {
			t := tables[j]
			contains, pos = t.probeForNumber(currentNumber)
			if contains {
				won = t.hasWon(pos.row, pos.column)
				if won {
					if len(tables) == 1 {
						return tables[0], currentNumber
					}
					tables = remove(tables, j)
					j--
				}
			}

		}
	}
	table := table{}
	return table, ""
}

func (t *table) probeForNumber(number string) (bool, position) {
	// checks if the table has new proposed number.
	// if yes it sets the respective successtable field to true and returns true to indicate a hit.
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if t.numbersTable[i][j] == number {
				t.successTable[i][j] = true
				return true, position{strconv.Itoa(i), strconv.Itoa(j)}
			}
		}
	}
	return false, position{}
}

func (t *table) hasWon(row string, column string) bool {
	// checks if the table is a winning table
	// checks lines that are affected by newest number
	r, _ := strconv.Atoi(row)
	c, _ := strconv.Atoi(column)
	// check row
	for i := 0; i < 5; i++ {
		if !t.successTable[r][i] {
			break
		}
		if i == 4 {
			return true
		}
	}
	// check column
	for j := 0; j < 5; j++ {
		if !t.successTable[j][c] {
			break
		}
		if j == 4 {
			return true
		}
	}
	return false
}

func (t *table) computeScore(winningNumber string) int {
	var score int
	w, _ := strconv.Atoi(winningNumber)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !t.successTable[i][j] {
				n, _ := strconv.Atoi(t.numbersTable[i][j])
				score += n
			}
		}
	}
	score *= w
	return score
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func remove(s []table, i int) []table {
	// Remove the element at index i from a.
	copy(s[i:], s[i+1:])  // Shift a[i+1:] left one index.
	s[len(s)-1] = table{} // Erase last element (write zero value).
	s = s[:len(s)-1]      // Truncate slice.
	return s
}
