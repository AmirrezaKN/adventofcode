package day01

import (
	"strconv"
	"strings"
)

func Question02_1(input []string) int {
	sum := 0

	games := helper(input)

	for _, game := range games {
		valid := true
		for _, round := range game.Rounds {
			for _, pair := range round.Pairs {
				// only 12 red cubes, 13 green cubes, and 14 blue cubes.
				if pair.Color == "red" && pair.QTY > 12 {
					valid = false
				}

				if pair.Color == "green" && pair.QTY > 13 {
					valid = false
				}

				if pair.Color == "blue" && pair.QTY > 14 {
					valid = false
				}
			}
		}

		if valid {
			sum += game.ID
		}
	}

	return sum
}

func Question02_2(input []string) int {
	sum := 0

	games := helper(input)

	for _, game := range games {
		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		for _, round := range game.Rounds {
			for _, pair := range round.Pairs {
				if pair.Color == "red" && pair.QTY > maxRed {
					maxRed = pair.QTY
				}

				if pair.Color == "green" && pair.QTY > maxGreen {
					maxGreen = pair.QTY
				}

				if pair.Color == "blue" && pair.QTY > maxBlue {
					maxBlue = pair.QTY
				}
			}
		}

		sum += (maxRed * maxBlue * maxGreen)
	}

	return sum
}

func helper(items []string) []game {
	res := []game{}

	for _, item := range items {
		g := game{
			Rounds: []round{},
		}
		// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		GameSlice := strings.Split(item, ": ")

		GameID, err := strconv.Atoi(strings.Split(GameSlice[0], " ")[1])
		if err != nil {
			panic(err)
		}

		g.ID = GameID

		roundsStrs := strings.Split(GameSlice[1], "; ")
		for _, roundStr := range roundsStrs {
			r := round{
				Pairs: []pair{},
			}
			roundSlice := strings.Split(roundStr, ", ")
			for _, pairStr := range roundSlice {
				p := pair{}
				pairSlice := strings.Split(pairStr, " ")

				p.Color = pairSlice[1]
				p.QTY, err = strconv.Atoi(pairSlice[0])
				if err != nil {
					panic(err)
				}

				r.Pairs = append(r.Pairs, p)

			}

			g.Rounds = append(g.Rounds, r)
		}

		res = append(res, g)
	}

	return res
}

type game struct {
	ID     int
	Rounds []round
}

type round struct {
	Pairs []pair
}

type pair struct {
	Color string
	QTY   int
}
