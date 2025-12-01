package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	Left Direction = iota + 1
	Right
)

type Dial struct {
	minNumber int
	maxNumber int
	current   int
	password  int
}

func NewDial(min, max, current int) *Dial {
	return &Dial{
		minNumber: min,
		maxNumber: max,
		current:   current,
		password:  0,
	}
}

func (d *Dial) Rotate2(direction Direction, distance int) {
	prevPos := d.current

	// Count complete 100-step loops (each passes through 0 once)
	d.password += distance / 100

	// Apply only the remaining distance after complete loops
	remainder := distance % 100

	switch direction {
	case Left:
		d.current -= remainder
		if d.current < 0 {
			d.current += 100
			// We crossed 0 if we didn't start or land on 0
			if d.current != 0 && prevPos != 0 {
				d.password++
			}
		}
	case Right:
		d.current += remainder
		if d.current > 99 {
			d.current -= 100
			// We crossed 0 if we didn't start or land on 0
			if d.current != 0 && prevPos != 0 {
				d.password++
			}
		}
	}

	// Always count landing on 0
	if d.current == 0 {
		d.password++
	}
}

func (d *Dial) Rotate1(direction Direction, distance int) {
	switch direction {
	case Left:
		d.current = (d.current - distance) % 100
		if d.current < 0 {
			d.current += 100
		}
	case Right:
		d.current = (d.current + distance) % 100
	}

	if d.current == 0 {
		d.password++
	}
}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	dial := NewDial(0, 99, 50)

	instructions := strings.SplitSeq(string(bytes), "\n")

	for instruction := range instructions {
		if instruction == "" {
			continue
		}

		direction, distance, err := parse(instruction)
		if err != nil {
			panic(err)
		}

		dial.Rotate2(direction, distance)
	}

	fmt.Printf("Password: %d\n", dial.password)
}

func parse(instruction string) (Direction, int, error) {
	var direction Direction

	switch instruction[0] {
	case 'L':
		direction = Left
	case 'R':
		direction = Right
	default:
		return 0, 0, fmt.Errorf("invalid direction: %s", string(instruction[0]))
	}

	distance, err := strconv.ParseInt(instruction[1:], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("could not parse distance: %w", err)
	}

	return direction, int(distance), nil
}
