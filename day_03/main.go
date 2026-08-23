package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partOne(input string) {
	wires := strings.Split(input, "\n")
	firstWire := strings.Split(wires[0], ",")
	secondWire := strings.Split(wires[1], ",")

	type Coordinate struct {
		X int
		Y int
	}

	curr := Coordinate{0, 0}
	seen := make(map[Coordinate]int)

	for _, val := range firstWire {
		direction := val[0]
		amount, _ := strconv.Atoi(val[1:])

		switch direction {
		case 'R':
			for range amount {
				curr.X += 1
				seen[curr] += 1
			}
		case 'L':
			for range amount {
				curr.X -= 1
				seen[curr] += 1
			}
		case 'U':
			for range amount {
				curr.Y -= 1
				seen[curr] += 1
			}
		case 'D':
			for range amount {
				curr.Y += 1
				seen[curr] += 1
			}
		}
	}

	// I'm lazy, this ensures it's at most 1.
	for coord := range seen {
		seen[coord] = 1
	}

	curr = Coordinate{0, 0}

	for _, val := range secondWire {
		direction := val[0]
		amount, _ := strconv.Atoi(val[1:])

		switch direction {
		case 'R':
			for range amount {
				curr.X += 1
				seen[curr] += 1
			}
		case 'L':
			for range amount {
				curr.X -= 1
				seen[curr] += 1
			}
		case 'U':
			for range amount {
				curr.Y -= 1
				seen[curr] += 1
			}
		case 'D':
			for range amount {
				curr.Y += 1
				seen[curr] += 1
			}
		}
	}

	smallestDistance := 1000000000000000
	for coord, val := range seen {
		if val > 1 {
			distance := max(coord.X, -coord.X) + max(coord.Y, -coord.Y)
			// fmt.Printf("coord x: %d, coord y: %d, val: %d -> %d\n", coord.X, coord.Y, val, distance)

			if distance < smallestDistance {
				smallestDistance = distance
			}
		}
	}

	fmt.Printf("Part 1: %d\n", smallestDistance)
}

func partTwo(input string) {
	wires := strings.Split(input, "\n")
	firstWire := strings.Split(wires[0], ",")
	secondWire := strings.Split(wires[1], ",")

	type Coordinate struct {
		X int
		Y int
	}

	curr := Coordinate{0, 0}
	stepTrackerOne := make(map[Coordinate]int)
	steps := 0

	for _, val := range firstWire {
		direction := val[0]
		amount, _ := strconv.Atoi(val[1:])

		switch direction {
		case 'R':
			for range amount {
				curr.X += 1
				steps += 1
				stepTrackerOne[curr] = steps
			}
		case 'L':
			for range amount {
				curr.X -= 1
				steps += 1
				stepTrackerOne[curr] = steps
			}
		case 'U':
			for range amount {
				curr.Y -= 1
				steps += 1
				stepTrackerOne[curr] = steps
			}
		case 'D':
			for range amount {
				curr.Y += 1
				steps += 1
				stepTrackerOne[curr] = steps
			}
		}
	}

	stepTrackerTwo := make(map[Coordinate]int)
	curr = Coordinate{0, 0}
	steps = 0

	for _, val := range secondWire {
		direction := val[0]
		amount, _ := strconv.Atoi(val[1:])

		switch direction {
		case 'R':
			for range amount {
				curr.X += 1
				steps += 1
				stepTrackerTwo[curr] = steps
			}
		case 'L':
			for range amount {
				curr.X -= 1
				steps += 1
				stepTrackerTwo[curr] = steps
			}
		case 'U':
			for range amount {
				curr.Y -= 1
				steps += 1
				stepTrackerTwo[curr] = steps
			}
		case 'D':
			for range amount {
				curr.Y += 1
				steps += 1

				stepTrackerTwo[curr] = steps
			}
		}
	}

	smallestNumOfSteps := 1000000000000000000
	for coord, stepOne := range stepTrackerOne {
		if stepTrackerTwo[coord] > 0 {
			numberOfSteps := stepOne + stepTrackerTwo[coord]
			// fmt.Printf("coord x: %d, coord y: %d -> %d\n", coord.X, coord.Y, numberOfSteps)

			if numberOfSteps < smallestNumOfSteps {
				smallestNumOfSteps = numberOfSteps
			}
		}

	}

	fmt.Printf("Part 2: %d\n", smallestNumOfSteps)
}

func main() {
	inputFile := "input.txt"

	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}

	inputBytes, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("couldn't read file: %s\n", err)
		return
	}
	input := string(inputBytes)

	partOne(input)
	partTwo(input)
}
