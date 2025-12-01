package day01

import (
	"strconv"
)

func Question01_1(input []string) int {
	sum := 0
	min := 0
	max := 99
	curr := 50

	for _, item := range input {
		if item == "" {
			continue
		}

		rotation := string(item[0])
		distance, _ := strconv.Atoi(string(item[1:]))

		if distance >= 100 {
			distance = distance % 100
		}

		if distance == 0 {
			continue
		}

		if rotation == "L" {
			distance = -distance
		}

		if curr+distance > max {
			curr = curr + distance - max - 1
		} else if curr+distance < min {
			curr = max + (curr + distance) + 1
		} else {
			curr += distance
		}

		if curr == 0 {
			sum++
		}

	}

	return sum
}

func Question01_2(input []string) int {
	sum := 0
	min := 0
	max := 99
	curr := 50

	for _, item := range input {
		if item == "" {
			continue
		}

		rotation := string(item[0])
		distance, _ := strconv.Atoi(string(item[1:]))

		if distance >= 100 {
			sum += distance / 100
			distance = distance % 100
		}

		if distance == 0 {
			continue
		}

		if rotation == "L" {
			distance = -distance
		}

		if curr+distance > max {
			flag := curr == 0
			curr = curr + distance - max - 1
			if curr != 0 && !flag {
				sum++
			}
		} else if curr+distance < min {
			flag := curr == 0
			curr = max + (curr + distance) + 1
			if curr != 0 && !flag {
				sum++
			}
		} else {
			curr += distance
		}

		if curr == 0 {
			sum++
		}
	}

	return sum
}
