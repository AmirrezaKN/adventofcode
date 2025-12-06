package day05

import (
	"strconv"
	"strings"
)

func Question05_1(input []string) int {
	sum := 0
	ranges := make([][2]int, 0)

	for i := range input {
		if input[i] == "" {
			input = input[i+1:]
			break
		}

		min, _ := strconv.Atoi(strings.Split(input[i], "-")[0])
		max, _ := strconv.Atoi(strings.Split(input[i], "-")[1])
		ranges = append(ranges, [2]int{min, max})
	}

	for _, item := range input {
		if item == "" {
			continue
		}

		num, _ := strconv.Atoi(item)

		for _, r := range ranges {
			if num >= r[0] && num <= r[1] {
				sum++
				break
			}
		}
	}

	return sum
}

func Question05_2(input []string) uint {
	var sum uint = 0
	ranges := make([][2]uint, 0)

	for i := range input {
		if input[i] == "" {
			input = input[i+1:]
			break
		}

		min, _ := strconv.Atoi(strings.Split(input[i], "-")[0])
		max, _ := strconv.Atoi(strings.Split(input[i], "-")[1])
		ranges = append(ranges, [2]uint{uint(min), uint(max)})
	}

	for {
		newRanges, flag := Combine(ranges)
		if !flag {
			break
		}
		ranges = newRanges
	}

	for _, r := range ranges {
		sum += r[1] - r[0] + 1
	}

	return sum
}

func Combine(ranges [][2]uint) ([][2]uint, bool) {
	newRanges := [][2]uint{}
	flag := false

	for i, r1 := range ranges {
		for j, r2 := range ranges {
			if r1[0] <= r2[0] && r1[1] >= r2[1] {
				continue
			}

			if r1[0] <= r2[0] && r1[1] >= r2[0] {
				newRanges = append(newRanges, [2]uint{r1[0], r2[1]})
				for k, item := range ranges {
					if k != i && k != j {
						newRanges = append(newRanges, item)
					}
				}

				return newRanges, true
			}
			if r1[0] <= r2[1] && r1[1] >= r2[1] {
				newRanges = append(newRanges, [2]uint{r2[0], r1[1]})
				for k, item := range ranges {
					if k != i && k != j {
						newRanges = append(newRanges, item)
					}
				}

				return newRanges, true
			}
			if r1[0] >= r2[0] && r1[1] <= r2[1] {
				newRanges = append(newRanges, [2]uint{r1[0], r1[1]})
				for k, item := range ranges {
					if k != i && k != j {
						newRanges = append(newRanges, item)
					}
				}

				return newRanges, true
			}
		}
	}

	return ranges, flag
}
