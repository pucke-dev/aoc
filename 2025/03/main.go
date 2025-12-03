package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	banks := strings.SplitSeq(strings.TrimSpace(string(bytes)), "\n")

	totalJoltage := 0

	for bank := range banks {
		batteries := toBatteries(bank)
		fmt.Println("Batteries: ", batteries)

		subSequence := findLargestSubSequence(batteries, 12)
		fmt.Println("Largest Sub Sequence: ", subSequence)

		ssStrings := make([]string, 0, len(subSequence))
		for _, v := range subSequence {
			ssStrings = append(ssStrings, strconv.Itoa(v))
		}

		bankJoltage, _ := strconv.ParseInt(strings.Join(ssStrings, ""), 10, 64)
		totalJoltage += int(bankJoltage)
	}

	fmt.Println("Total Joltage: ", totalJoltage)
}

func findLargestSubSequence(sequence []int, length int) []int {
	result := make([]int, 0, length)

	start := 0
	for i := range length {
		// We have to select (length - i) more digits total
		// We must leave at least (length - i - 1) digits after our choice
		// So we can look from start to len(sequence) - (length - i - 1).
		end := len(sequence) - (length - i - 1)

		peekableSlice := sequence[start:end]

		maxDigit := slices.Max(peekableSlice)
		maxDigitIndex := slices.Index(peekableSlice, maxDigit)

		result = append(result, maxDigit)
		start += maxDigitIndex + 1
	}

	return result
}

func toBatteries(bank string) []int {
	batteries := make([]int, 0, len(bank))

	for _, battery := range bank {
		b, _ := strconv.ParseInt(string(battery), 10, 64)
		batteries = append(batteries, int(b))
	}

	return batteries
}
