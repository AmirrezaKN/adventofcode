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
