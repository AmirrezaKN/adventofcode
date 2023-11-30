package day01

import "strconv"

func Question01_1(input []string) int {
	max, temp := 0, 0

	for _, item := range input {
		if item != "" {
			number, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			temp = temp + number
			continue
		}

		if temp > max {
			max = temp
		}
		temp = 0
	}

	if temp > max {
		max = temp
	}

	return max
}

func Question01_2(input []string) int {
	temp := 0
	max := []int{0}

	for _, item := range input {
		if item != "" {
			number, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			temp = temp + number
			continue
		}

		if temp > max[len(max)-1] {
			max = append(max, temp)
		}
		temp = 0
	}

	if temp > max[len(max)-1] {
		max = append(max, temp)
	}

	return max[len(max)-1] + max[len(max)-2] + max[len(max)-3]
}
