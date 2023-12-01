package day01

import (
	"strconv"
)

func Question01_1(input []string) int {
	sum := 0

	for _, item := range input {
		first, last := 0, 0
		firstSeen := false

		for _, r := range item {
			num, err := strconv.Atoi(string(r))
			if err != nil {
				continue
			}

			if !firstSeen {
				first = num
				firstSeen = true
			}

			last = num
		}
		sum += (first * 10) + last
	}

	return sum
}

func Question01_2(input []string) int {
	sum := 0
	// one, two, three, four, five, six, seven, eight, and nine
	for _, item := range input {
		first, last := 0, 0
		firstSeen := false

		for i := range item {
			switch string(item[i]) {
			// one
			case "o":
				helper(item, "one", i, 1, &first, &last, &firstSeen)
			// eight
			case "e":
				helper(item, "eight", i, 8, &first, &last, &firstSeen)
			// four, five
			case "f":
				helper(item, "four", i, 4, &first, &last, &firstSeen)

				helper(item, "five", i, 5, &first, &last, &firstSeen)
			// six, seven
			case "s":
				helper(item, "six", i, 6, &first, &last, &firstSeen)

				helper(item, "seven", i, 7, &first, &last, &firstSeen)
			// nine
			case "n":
				helper(item, "nine", i, 9, &first, &last, &firstSeen)
			// two, three
			case "t":
				helper(item, "two", i, 2, &first, &last, &firstSeen)
				helper(item, "three", i, 3, &first, &last, &firstSeen)

			default:
				num, err := strconv.Atoi(string(item[i]))
				if err != nil {
					continue
				}

				if !firstSeen {
					first = num
					firstSeen = true
				}
				last = num
			}
		}
		sum += (first * 10) + last
	}

	return sum
}

func helper(item string, word string, i, number int, first, last *int, firstSeen *bool) {
	if len(word) > len(item)-i {
		return
	}

	if item[i:i+len(word)] == word {
		if firstSeen != nil && !*firstSeen {
			*first = number
			*firstSeen = true
		}
		*last = number
	}
}
