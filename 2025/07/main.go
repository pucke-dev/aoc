package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Total splits: ", countSplits(bytes))
	fmt.Println("Total timelines: ", countTimelines(bytes))
}

func countTimelines(bytes []byte) int {
	grid := NewGrid(bytes)
	memo := make(map[string]int)

	return dfs(grid, 0, slices.Index(grid[0], "S"), memo)
}

func countSplits(bytes []byte) int {
	grid := NewGrid(bytes)

	splits := 0

	for row := range grid {
		for col := range grid[row] {
			char := grid[row][col]

			switch char {
			case "S":
				addBeam(grid, row+1, col)
			case "|":
				peek := peekBelow(grid, row, col)
				switch peek {
				case ".":
					addBeam(grid, row+1, col)
				case "^":
					addBeam(grid, row+1, col-1)
					addBeam(grid, row+1, col+1)
					splits++
				}

			}
		}
	}

	return splits
}

func peekBelow(grid [][]string, row, col int) any {
	if row >= len(grid)-1 {
		return ""
	}

	return grid[row+1][col]
}

func addBeam(grid [][]string, row, col int) {
	grid[row][col] = "|"
}

func NewGrid(bytes []byte) [][]string {
	grid := make([][]string, 0)

	for row := range strings.SplitSeq(strings.TrimSpace(string(bytes)), "\n") {
		r := make([]string, 0)
		r = append(r, strings.Split(strings.TrimSuffix(row, "\n"), "")...)
		grid = append(grid, r)
	}

	return grid
}

func dfs(grid [][]string, row, col int, memo map[string]int) int {
	if row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return 1
	}
	key := fmt.Sprintf("%d,%d", row, col)
	if v, ok := memo[key]; ok {
		return v
	}

	result := 0
	switch grid[row][col] {
	case "S", ".":
		result = dfs(grid, row+1, col, memo)

	case "^":
		result = dfs(grid, row+1, col-1, memo) + dfs(grid, row+1, col+1, memo)
	}

	memo[key] = result
	return result
}
