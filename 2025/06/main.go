package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Grand total: ", part1(bytes))
	fmt.Println("Part 2: ", part2(bytes))
}

func part2(bytes []byte) int {
	grid := make([][]string, 0)

	for row := range strings.SplitSeq(string(bytes), "\n") {
		r := make([]string, 0)
		r = append(r, strings.Split(strings.TrimSuffix(row, "\n"), "")...)

		if len(r) == 0 {
			continue
		}

		grid = append(grid, r)
	}

	// printGrid(grid)
	// fmt.Println("Rows: ", len(grid))
	// fmt.Println("Cols: ", len(grid[0]))

	total := 0
	operands := make([]int, 0)
	for col := len(grid[0]) - 1; col >= 0; col-- {
		charsInColumn := make([]string, len(grid))

		for row := 0; row < len(grid)-1; row++ {
			charsInColumn = append(charsInColumn, grid[row][col])
		}

		// fmt.Println("Chars in col: ", charsInColumn)

		number, _ := strconv.Atoi(strings.TrimSpace(strings.Join(charsInColumn, "")))
		if number == 0 {
			continue
		}

		operands = append(operands, number)

		// fmt.Println("operands: ", operands)

		switch grid[len(grid)-1][col] {
		case "":
			continue
		case "*":
			total += multiplyOperands(operands)
			operands = []int{}

		case "+":
			total += sumOperands(operands)
			operands = []int{}
		}
	}

	return total
}

func sumOperands(operands []int) int {
	total := 0
	for _, v := range operands {
		total += v
	}

	return total
}

func multiplyOperands(operands []int) int {
	total := 1
	for _, v := range operands {
		total *= v
	}

	return total
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Printf("%v", row)
		fmt.Printf("\n")
	}
}

func part1(bytes []byte) int {
	grid := make([][]string, 0)

	for row := range strings.SplitSeq(strings.TrimSpace(string(bytes)), "\n") {
		r := make([]string, 0)
		r = append(r, strings.Fields(row)...)
		grid = append(grid, r)
	}

	total := 0

	for col := 0; col < len(grid[0]); col++ {
		operation := grid[len(grid)-1][col]

		columnTotal := 0
		for row := 0; row < len(grid)-1; row++ {
			v, _ := strconv.Atoi(grid[row][col])

			switch operation {
			case "+":
				columnTotal += v
			case "*":
				if columnTotal == 0 {
					columnTotal++
				}

				columnTotal *= v
			}
		}

		total += columnTotal
	}

	return total
}
