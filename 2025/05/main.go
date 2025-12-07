package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	From int
	To   int
}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	ingredientDB := strings.Split(string(bytes), "\n\n")
	freshRanges := ingredientDB[0]
	availableIDs := ingredientDB[1]

	freshIDMap := newFreshMap(strings.Split(freshRanges, "\n"))

	fmt.Println("Available fresh ids: ", countFreshIDs(freshIDMap, strings.Split(availableIDs, "\n")))
	fmt.Println("Number of all freshIDs: ", countAllFreshIDs(freshIDMap))
}

func countAllFreshIDs(freshIDMap map[Range]bool) int {
	ranges := slices.Collect(maps.Keys(freshIDMap))

	byRangeStart := func(a, b Range) int {
		return a.From - b.From
	}
	slices.SortFunc(ranges, byRangeStart)

	merged := []*Range{&ranges[0]}

	for i := 1; i < len(ranges); i++ {
		current := ranges[i]
		last := merged[len(merged)-1]

		if current.From <= last.To+1 {
			if current.To > last.To {
				last.To = current.To
			}
		} else {
			merged = append(merged, &current)
		}
	}

	total := 0
	for _, r := range merged {
		total += r.To - r.From + 1
	}

	return total
}

func newFreshMap(freshRanges []string) map[Range]bool {
	freshMap := make(map[Range]bool, len(freshRanges))

	for _, r := range freshRanges {
		tmp := strings.Split(r, "-")
		start, _ := strconv.ParseInt(tmp[0], 10, 64)
		end, _ := strconv.ParseInt(tmp[1], 10, 64)

		freshMap[Range{int(start), int(end)}] = true
	}

	return freshMap

}

func countFreshIDs(freshMap map[Range]bool, ids []string) int {
	count := 0

	for _, i := range ids {
		id, _ := strconv.ParseInt(i, 10, 64)

		for r := range freshMap {
			if int64(r.From) <= id && id <= int64(r.To) {
				count++
				break
			}
		}
	}

	return count
}
