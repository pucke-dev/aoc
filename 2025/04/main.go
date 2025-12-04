package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	xDimension = 0
	yDimension = 0
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	grid := NewGrid(string(bytes))

	xDimension = len(grid)
	yDimension = len(grid[0])

	accessableForklifts := 0

	for finished := false; !finished; {

		removedRolls := 0

		for x := 0; x < xDimension; x++ {
			for y := 0; y < yDimension; y++ {
				adjacents := 0

				if grid[x][y] == "." {
					continue
				}

				if hasTopLeft(grid, x, y) {
					adjacents++
				}

				if hasTop(grid, x, y) {
					adjacents++
				}

				if hasTopRight(grid, x, y) {
					adjacents++
				}

				if hasRight(grid, x, y) {
					adjacents++
				}

				if hasBottomRight(grid, x, y) {
					adjacents++
				}

				if hasBottom(grid, x, y) {
					adjacents++
				}

				if hasBottomLeft(grid, x, y) {
					adjacents++
				}

				if hasLeft(grid, x, y) {
					adjacents++
				}

				if adjacents < 4 {
					accessableForklifts++

					// Part 2: Remove accessible
					grid[x][y] = "."
					removedRolls++
				}
			}
		}

		if removedRolls == 0 {
			finished = true
		}
	}

	fmt.Println("Accessable Forklifts: ", accessableForklifts)
}

// o-o-o y
// x-P-o
// o-o-o
// x
func hasLeft(grid [][]string, x, y int) bool {
	if y-1 < 0 {
		return false
	}
	return grid[x][y-1] == "@"
}

// o-o-o y
// o-P-o
// x-o-o
// x
func hasBottomLeft(grid [][]string, x, y int) bool {
	if x+1 >= xDimension || y-1 < 0 {
		return false
	}
	return grid[x+1][y-1] == "@"
}

// o-o-o y
// o-P-o
// o-x-o
// x
func hasBottom(grid [][]string, x, y int) bool {
	if x+1 >= xDimension {
		return false
	}
	return grid[x+1][y] == "@"
}

// o-o-o y
// o-P-o
// o-o-x
// x
func hasBottomRight(grid [][]string, x, y int) bool {
	if x+1 >= xDimension || y+1 >= yDimension {
		return false
	}
	return grid[x+1][y+1] == "@"
}

// o-o-o y
// o-P-x
// o-o-o
// x
func hasRight(grid [][]string, x, y int) bool {
	if y+1 >= yDimension {
		return false
	}
	return grid[x][y+1] == "@"
}

// o-o-x y
// o-P-o
// o-o-o
// x
func hasTopRight(grid [][]string, x, y int) bool {
	if x-1 < 0 || y+1 >= yDimension {
		return false
	}
	return grid[x-1][y+1] == "@"
}

// o-x-o y
// o-P-o
// o-o-o
// x
func hasTop(grid [][]string, x, y int) bool {
	if x-1 < 0 {
		return false
	}
	return grid[x-1][y] == "@"
}

// x-o-o y
// o-P-o
// o-o-o
// x
func hasTopLeft(grid [][]string, x, y int) bool {
	if x-1 < 0 || y-1 < 0 {
		return false
	}
	return grid[x-1][y-1] == "@"
}

func NewGrid(input string) [][]string {
	rows := strings.Split(strings.TrimSpace(input), "\n")

	grid := make([][]string, 0, len(rows))

	for _, row := range rows {
		cols := make([]string, 0, len(row))

		for _, col := range row {
			cols = append(cols, string(col))
		}

		grid = append(grid, cols)
	}

	return grid
}
