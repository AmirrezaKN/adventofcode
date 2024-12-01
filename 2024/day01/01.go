package day01

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func Question01_1(input []string) int {
	sum := 0
	left, right := []int{}, []int{}

	for _, item := range input {
		items := strings.Split(item, "   ")

		leftInt, err := strconv.Atoi(items[0])
		if err != nil {
			panic(err)
		}
		left = append(left, leftInt)

		rightInt, err := strconv.Atoi(items[1])
		if err != nil {
			panic(err)
		}
		right = append(right, rightInt)
	}

	sort.Slice(left, func(i, j int) bool {
		return left[j] < left[i]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[j] < right[i]
	})

	for i := 0; i < len(input); i++ {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}

	return sum
}

func Question01_2(input []string) int {
	sum := 0
	left, right := []int{}, []int{}

	for _, item := range input {
		items := strings.Split(item, "   ")

		leftInt, err := strconv.Atoi(items[0])
		if err != nil {
			panic(err)
		}
		left = append(left, leftInt)

		rightInt, err := strconv.Atoi(items[1])
		if err != nil {
			panic(err)
		}
		right = append(right, rightInt)
	}

	sort.Slice(left, func(i, j int) bool {
		return left[j] < left[i]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[j] < right[i]
	})

	rightMap := make(map[int]int)

	for _, item := range right {
		rightMap[item] = 0
	}

	for _, item := range right {
		if t, ok := rightMap[item]; ok {
			rightMap[item] = t + 1
		}
	}

	for _, item := range left {
		if t, ok := rightMap[item]; ok {
			sum += item * t
		}
	}

	return sum
}
