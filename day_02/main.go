package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func runProgram(inputs []int) {
	for i := 0; i < len(inputs); i += 4 {
		opcode := inputs[i]
		if opcode == 99 {
			break
		}

		lhs := inputs[i+1]
		rhs := inputs[i+2]
		destination := inputs[i+3]

		// Lazy way to catch invalid inputs.
		if destination >= len(inputs) || lhs >= len(inputs) || rhs >= len(inputs) {
			break
		}

		switch opcode {
		case 1:
			inputs[destination] = inputs[lhs] + inputs[rhs]
		case 2:
			inputs[destination] = inputs[lhs] * inputs[rhs]
		}
	}
}

func partOne(input string, useProgramAlarm bool) {
	inputStrs := strings.Split(input, ",")
	inputs := []int{}

	for _, x := range inputStrs {
		num, _ := strconv.Atoi(x)
		inputs = append(inputs, num)
	}

	if useProgramAlarm {
		// For part 1, we have the "1202 program alarm".
		inputs[1] = 12
		inputs[2] = 2
	}

	runProgram(inputs)

	fmt.Println("Part 1:", inputs[0])
}

func partTwo(input string) {
	inputStrs := strings.Split(input, ",")
	originalInputs := []int{}

	for _, x := range inputStrs {
		num, _ := strconv.Atoi(x)
		originalInputs = append(originalInputs, num)
	}

	for i := range 1000 {
		for j := range 1000 {
			noun := i
			verb := j

			inputs := make([]int, len(originalInputs))
			copy(inputs, originalInputs)

			inputs[1] = noun
			inputs[2] = verb

			runProgram(inputs)

			if inputs[0] == 19690720 {
				fmt.Println("Part 2:", 100*noun+verb)
				break
			}
		}
	}
}

func main() {
	inputFile := "input.txt"

	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}

	inputBytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("couldn't read file: %s", err)
	}
	input := string(inputBytes)

	partOne(input, strings.HasSuffix(input, "input.txt"))
	partTwo(input)
}
