package day01

import "strconv"

func Question03_1(input []string) int {
	sum := 0

	for _, item := range input {
		if item == "" {
			continue
		}

		max := 0

		leftIndex := 0
		rightIndex := len(item) - 1

		for rightIndex > leftIndex {
			leftPart := item[leftIndex:rightIndex]
			rightPart := item[rightIndex:]
			temp := (findMax(leftPart) * 10) + findMax(rightPart)

			if temp > max {
				max = temp
			} else {
				rightIndex--
			}
		}

		sum += max
	}

	return sum
}

func findMax(input string) int {
	max := 0
	for _, item := range input {
		num, _ := strconv.Atoi(string(item))
		if num > max {
			max = num
		}
	}
	return max
}

func Question03_2(input []string) int {
	sum := 0

	for _, item := range input {
		if item == "" {
			continue
		}

		maxStr := ""
		leftIndex := 0
		rightIndex := len(item) - 12
		previousIndex := 0

		for rightIndex <= len(item)-1 {
			activePart := item[leftIndex : rightIndex+1]

			activeMax, index := findMaxWithIndex(activePart)

			maxStr += strconv.Itoa(activeMax)
			previousIndex = leftIndex
			leftIndex = previousIndex + index + 1
			rightIndex++
		}

		max, _ := strconv.Atoi(maxStr)
		sum += max
	}

	return sum
}

func findMaxWithIndex(input string) (int, int) {
	max := 0
	index := 0
	for i, item := range input {
		num, _ := strconv.Atoi(string(item))
		if num > max {
			max = num
			index = i
		}
	}
	return max, index
}
