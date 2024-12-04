package day01

import (
	"math"
	"strconv"
	"strings"
)

func Part1(input []string) int {
	sum := 0
	for _, line := range input {
		numbers := strings.Split(line, " ")

		if numbers[0] == numbers[1] {
			continue
		}

		firstNumber, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}

		secondNumber, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}

		desc := secondNumber < firstNumber

		i := 0
		j := 1
		flag := false
		for {
			if j >= len(numbers) {
				break
			}

			iNumber, err := strconv.Atoi(numbers[i])
			if err != nil {
				panic(err)
			}

			jNumber, err := strconv.Atoi(numbers[j])
			if err != nil {
				panic(err)
			}

			if iNumber == jNumber {
				flag = true
				break
			}

			if desc != (jNumber < iNumber) {
				flag = true
				break
			}

			if math.Abs(float64(jNumber-iNumber)) > 3 {
				flag = true
				break
			}

			i = j
			j++
		}

		if !flag {
			sum++
		}

	}

	return sum
}

func Part2(input []string) int {
	sum := 0
	for _, line := range input {
		strs := strings.Split(line, " ")

		numbers := []int{}
		for _, str := range strs {
			number, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}

			numbers = append(numbers, number)

		}

		flag := calc(numbers, 0)
		if flag {
			sum++
		}
	}

	return sum
}

func calc(numbers []int, round int) bool {
	if round > 1 {
		return false
	}

	if numbers[0] == numbers[1] {
		lol := make([]int, len(numbers))
		copy(lol, numbers)
		lol2 := make([]int, len(numbers))
		copy(lol2, numbers)

		t11 := numbers[:0]
		t12 := numbers[0+1:]
		temp1 := append(t11, t12...)

		t21 := lol[:1]
		t22 := lol[1+1:]
		temp2 := append(t21, t22...)

		return calc(temp1, round+1) || calc(temp2, round+1) || calc(lol2[1:], round+1)
	}

	firstNumber := numbers[0]
	secondNumber := numbers[1]

	desc := secondNumber < firstNumber

	i := 0
	j := 1
	flag := true
	for {
		if j >= len(numbers) {
			return flag
		}

		iNumber := numbers[i]
		jNumber := numbers[j]

		if iNumber == jNumber {
			lol := make([]int, len(numbers))
			copy(lol, numbers)
			lol2 := make([]int, len(numbers))
			copy(lol2, numbers)

			t11 := numbers[:i]
			t12 := numbers[i+1:]
			temp1 := append(t11, t12...)

			t21 := lol[:j]
			t22 := lol[j+1:]
			temp2 := append(t21, t22...)

			return calc(temp1, round+1) || calc(temp2, round+1) || calc(lol2[1:], round+1)

		}

		if desc != (jNumber < iNumber) {
			lol := make([]int, len(numbers))
			copy(lol, numbers)
			lol2 := make([]int, len(numbers))
			copy(lol2, numbers)

			t11 := numbers[:i]
			t12 := numbers[i+1:]
			temp1 := append(t11, t12...)

			t21 := lol[:j]
			t22 := lol[j+1:]
			temp2 := append(t21, t22...)

			return calc(temp1, round+1) || calc(temp2, round+1) || calc(lol2[1:], round+1)

		}

		if math.Abs(float64(jNumber-iNumber)) > 3 {
			lol := make([]int, len(numbers))
			copy(lol, numbers)
			lol2 := make([]int, len(numbers))
			copy(lol2, numbers)

			t11 := numbers[:i]
			t12 := numbers[i+1:]
			temp1 := append(t11, t12...)

			t21 := lol[:j]
			t22 := lol[j+1:]
			temp2 := append(t21, t22...)

			return calc(temp1, round+1) || calc(temp2, round+1) || calc(lol2[1:], round+1)

		}

		i = j
		j++
	}
}
