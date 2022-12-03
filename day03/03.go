package day03

import (
	"unicode"
)

func Question02_1(input []string) int {
	total := 0

	for _, str := range input {
		counts := make(map[rune]int)

		for _, r := range str[:len(str)/2] {
			counts[r] += 1
		}

		for _, r := range str[len(str)/2:] {
			if _, ok := counts[r]; ok {
				if unicode.IsLower(r) {
					total += int(r) - 96
				} else {
					total += int(r) - 38
				}
				delete(counts, r)
			}
		}
	}

	return total
}

func Question02_2(input []string) int {
	total := 0

	for i := 0; i < len(input); i += 3 {
		if i+2 > len(input) {
			break
		}

		seenMap := make(map[rune][]bool)
		for _, r := range input[i] {
			if len(seenMap[r]) == 0 {
				seenMap[r] = append(seenMap[r], true)
			}
		}

		for _, r := range input[i+1] {
			if len(seenMap[r]) == 1 {
				seenMap[r] = append(seenMap[r], true)
			}
		}

		for _, r := range input[i+2] {
			if len(seenMap[r]) == 2 {
				if unicode.IsLower(r) {
					total += int(r) - 96
				} else {
					total += int(r) - 38
				}
				break
			}
		}
	}

	return total
}
