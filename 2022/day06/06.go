package day06

import "math"

func Question06_1(str string) int {
	max := 0
	left := 0
	seen := make(map[rune]bool)

	for right, r := range str {
		for seen[r] {
			delete(seen, rune(str[left]))
			left++
		}
		seen[r] = true
		max = int(math.Max(float64(len(seen)), float64(max)))
		if max == 4 {
			return right + 1
		}
	}
	return max
}

func Question06_2(str string) int {
	max := 0
	left := 0
	seen := make(map[rune]bool)

	for right, r := range str {
		for seen[r] {
			delete(seen, rune(str[left]))
			left++
		}
		seen[r] = true
		max = int(math.Max(float64(len(seen)), float64(max)))
		if max == 14 {
			return right + 1
		}
	}
	return max
}
