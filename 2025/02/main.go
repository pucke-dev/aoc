package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var invalidIDMap = map[int64]int{}

func init() {
	invalidIDMap = generateAllInvalidIDs(10)
}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	ranges := strings.SplitSeq(strings.TrimSpace(string(bytes)), ",")

	var count int64
	var count2 int64

	for r := range ranges {
		ids := strings.Split(r, "-")

		firstID, _ := strconv.ParseInt(ids[0], 10, 64)
		lastID, _ := strconv.ParseInt(ids[1], 10, 64)

		for i := firstID; i <= lastID; i++ {
			if _, ok := invalidIDMap[i]; ok {
				count += i
			}

			if hasRepetitions(int(i)) {
				count2 += i
			}
		}
	}

	fmt.Printf("Part 1 Sum of invalid IDs: %d\n", count)
	fmt.Printf("Part 2 Sum of invalid IDs: %d\n", count2)
}

func hasRepetitions(i int) bool {
	s := strconv.Itoa(i)
	length := len(s)

	for patternLength := 1; patternLength <= length/2; patternLength++ {
		if length%patternLength != 0 {
			continue
		}

		pattern := s[:patternLength]
		repetitions := length / patternLength

		if strings.Repeat(pattern, repetitions) == s {
			return true
		}
	}

	return false
}

// generateAllInvalidIDs generates all IDs where the pattern is repeated exactly twice.
//
//	E.g. 123123
func generateAllInvalidIDs(digits int) map[int64]int {
	invalidIDs := make(map[int64]int)

	// Only even-length numbers can be invalid
	for totalDigits := 2; totalDigits <= digits; totalDigits += 2 {
		patternLength := totalDigits / 2
		multiplier := math.Pow(10, float64(patternLength)) + 1

		patternStart := math.Pow(10, float64(patternLength)-1)
		patternEnd := math.Pow(10, float64(patternLength)) - 1

		for p := patternStart; p <= patternEnd; p++ {
			invalidIDs[int64(p*multiplier)] = 1
		}
	}

	return invalidIDs
}
