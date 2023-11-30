package day05

import "fmt"

type move struct {
	amount      int
	source      int
	destination int
}

func Question05_1(stacks [9][]string, moves []move) string {
	result := ""

	for _, move := range moves {
		for i := 0; i < move.amount; i++ {
			cargo := stacks[move.source-1][len(stacks[move.source-1])-1]
			stacks[move.source-1] = stacks[move.source-1][:len(stacks[move.source-1])-1]
			stacks[move.destination-1] = append(stacks[move.destination-1], cargo)
		}
	}

	for _, stack := range stacks {
		result = fmt.Sprintf("%s%s", result, stack[len(stack)-1])
	}

	return result
}

func Question05_2(stacks [9][]string, moves []move) string {
	result := ""

	for _, move := range moves {
		cargo := []string{}
		for i := 0; i < move.amount; i++ {
			cargo = append(cargo, stacks[move.source-1][len(stacks[move.source-1])-1])
			stacks[move.source-1] = stacks[move.source-1][:len(stacks[move.source-1])-1]
		}

		for i := len(cargo) - 1; i >= 0; i-- {
			stacks[move.destination-1] = append(stacks[move.destination-1], cargo[i])
		}

	}

	for _, stack := range stacks {
		result = fmt.Sprintf("%s%s", result, stack[len(stack)-1])
	}

	return result
}
