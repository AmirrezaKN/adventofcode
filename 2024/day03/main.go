package day03

import (
	"regexp"
	"strconv"
	"strings"
)

func Part1(input []string) int {
	sum := 0
	for _, line := range input {
		// Compile the regex pattern
		regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

		// Find all matches
		matches := regex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])

			sum += a * b
		}
	}

	return sum
}

func Part2(input []string) int {
	sum := 0
	for _, line := range input {
		tokens := strings.Split(line, "don't()")
		finalLine := tokens[0]

		for i := 1; i < len(tokens); i++ {
			if index := strings.Index(tokens[i], "do()"); index != -1 {
				finalLine += tokens[i][index:]
			}
		}
		// Compile the regex pattern
		regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

		// Find all matches
		matches := regex.FindAllStringSubmatch(finalLine, -1)
		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])

			sum += a * b
		}
	}

	return sum
}
