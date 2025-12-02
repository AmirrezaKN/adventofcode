package day01

import (
	"fmt"
	"strconv"
	"strings"
)

func Question02_1(input []string) int {
	sum := 0

	for _, item := range input {
		if item == "" {
			continue
		}

		min, _ := strconv.Atoi(strings.Split(item, "-")[0])
		max, _ := strconv.Atoi(strings.Split(item, "-")[1])

		for i := min; i <= max; i++ {
			numStr := fmt.Sprint(i)

			if len(numStr)%2 != 0 {
				continue
			}

			if numStr[:len(numStr)/2] == numStr[len(numStr)/2:] {
				sum += i
			}
		}

	}

	return sum
}

func Question02_2(input []string) int {
	sum := 0

	for _, line := range input {
		if line == "" {
			continue
		}

		line = strings.ReplaceAll(line, "\n", "")
		min, _ := strconv.Atoi(strings.Split(line, "-")[0])
		max, _ := strconv.Atoi(strings.Split(line, "-")[1])

		for item := min; item <= max; item++ {
			numStr := fmt.Sprint(item)

			for batchSize := 1; batchSize <= len(numStr)/2; batchSize++ {
				if len(numStr)%batchSize != 0 {
					continue
				}

				firstStr := numStr[:batchSize]
				batchCount := len(numStr) / batchSize
				similarBatchCount := 1
				for i := batchSize; i <= len(numStr); i += batchSize {
					if i+batchSize > len(numStr) {
						temp := numStr[i:]
						if firstStr == temp {
							similarBatchCount++
							continue
						}
					} else {
						temp := numStr[i : i+batchSize]
						if firstStr == temp {
							similarBatchCount++
							continue
						}
					}
				}

				if batchCount == similarBatchCount {
					sum += item
					break
				}

			}
		}
	}

	return sum
}
