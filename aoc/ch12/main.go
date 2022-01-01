package main

import (
	"aoc.com/utils/dataStructures/graph"
	"aoc.com/utils/str"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	data, err := os.ReadFile("/home/jay/go/src/challenges/aoc/ch12/puzzleInput.txt")
	check(err)
	inputLines := strings.Split(string(data), "\n")
	g := graph.Graph{Nodes: make(map[string]*graph.Node)}

	// read in graph
	for _, line := range inputLines {
		srcId := strings.Split(line, "-")[0]
		destId := strings.Split(line, "-")[1]
		// add nodes if not present already
		g.AddEdge(srcId, destId, false)
	}
	paths := collectPaths(g)
	fmt.Println(len(paths))
}

func collectPaths(g graph.Graph) [][]string {
	startToEndPaths := make([][]string, 0)
	// start by the start node
	// add the start node to the todo list
	// pop element from todo list
	// go through neighbors and add a path to the todolist containing the path and the neighbor id added, but only if not two times a lowercase id is present.
	todoList := make([][]string, 0)
	todoList = append(todoList, []string{"start"})

	for len(todoList) != 0 {
		// pop off the top of the todoList
		currentPath := todoList[0]
		lastNodeId := currentPath[len(currentPath)-1]
		lastNode := g.GetNode(lastNodeId)
		if lastNodeId == "end" {
			startToEndPaths = append(startToEndPaths, currentPath)
			todoList = todoList[1:]
			continue
		}
		for _, node := range lastNode.GetConnections() {
			nodeId := node.GetId()
			nodeRune := []rune(node.GetId())
			if str.Contains(nodeId, currentPath) && ((unicode.IsLower(nodeRune[0]) && isDoubleLowerPath(currentPath)) || nodeId == "start" || nodeId == "end") {
				continue
			} else {
				// copy the path and append the neighbor to it
				pathCopy := make([]string, len(currentPath))
				copy(pathCopy, currentPath)
				pathCopy = append(pathCopy, string(nodeRune))
				// add the path to the todoList
				todoList = append(todoList, pathCopy)
			}
		}
		todoList = todoList[1:]
	}
	return startToEndPaths
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isDoubleLowerPath(path []string) bool {
	for checker := 0; checker < len(path)-1; checker++ {
		for contra := checker + 1; contra < len(path); contra++ {
			if path[checker] == path[contra] && unicode.IsLower([]rune(path[checker])[0]) {
				return true
			}
		}
	}
	return false
}
