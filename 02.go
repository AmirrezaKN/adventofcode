package main

type q2round struct {
	P1 string
	P2 string
}

func Question02_1(input []q2round) int {
	type moves struct {
		Rock     string
		Paper    string
		Scissors string
	}

	player1Moves := moves{
		Rock:     "A",
		Paper:    "B",
		Scissors: "C",
	}

	player2Moves := moves{
		Rock:     "X",
		Paper:    "Y",
		Scissors: "Z",
	}

	shapeScores := struct {
		Rock     int
		Paper    int
		Scissors int
	}{
		Rock:     1,
		Paper:    2,
		Scissors: 3,
	}

	resultScores := struct {
		Win  int
		Draw int
		Lose int
	}{
		Win:  6,
		Draw: 3,
		Lose: 0,
	}

	score := 0

	for _, item := range input {
		switch item.P2 {
		case player2Moves.Rock:
			switch item.P1 {
			case player1Moves.Rock:
				score = score + shapeScores.Rock + resultScores.Draw
			case player1Moves.Paper:
				score = score + shapeScores.Rock + resultScores.Lose
			case player1Moves.Scissors:
				score = score + shapeScores.Rock + resultScores.Win
			}
		case player2Moves.Paper:
			switch item.P1 {
			case player1Moves.Rock:
				score = score + shapeScores.Paper + resultScores.Win
			case player1Moves.Paper:
				score = score + shapeScores.Paper + resultScores.Draw
			case player1Moves.Scissors:
				score = score + shapeScores.Paper + resultScores.Lose
			}
		case player2Moves.Scissors:
			switch item.P1 {
			case player1Moves.Rock:
				score = score + shapeScores.Scissors + resultScores.Lose
			case player1Moves.Paper:
				score = score + shapeScores.Scissors + resultScores.Win
			case player1Moves.Scissors:
				score = score + shapeScores.Scissors + resultScores.Draw
			}
		}
	}

	return score
}

func Question02_2(input []q2round) int {
	type moves struct {
		Rock     string
		Paper    string
		Scissors string
	}

	player1Moves := moves{
		Rock:     "A",
		Paper:    "B",
		Scissors: "C",
	}

	player2Moves := struct {
		Win  string
		Draw string
		Lose string
	}{
		Win:  "Z",
		Draw: "Y",
		Lose: "X",
	}

	shapeScores := struct {
		Rock     int
		Paper    int
		Scissors int
	}{
		Rock:     1,
		Paper:    2,
		Scissors: 3,
	}

	resultScores := struct {
		Win  int
		Draw int
		Lose int
	}{
		Win:  6,
		Draw: 3,
		Lose: 0,
	}

	score := 0

	for _, item := range input {
		switch item.P2 {
		case player2Moves.Win:
			switch item.P1 {
			case player1Moves.Rock:
				score = score + shapeScores.Paper + resultScores.Win
			case player1Moves.Paper:
				score = score + shapeScores.Scissors + resultScores.Win
			case player1Moves.Scissors:
				score = score + shapeScores.Rock + resultScores.Win
			}
		case player2Moves.Lose:
			switch item.P1 {
			case player1Moves.Rock:
				score = score + shapeScores.Scissors + resultScores.Lose
			case player1Moves.Paper:
				score = score + shapeScores.Rock + resultScores.Lose
			case player1Moves.Scissors:
				score = score + shapeScores.Paper + resultScores.Lose
			}
		case player2Moves.Draw:
			switch item.P1 {
			case player1Moves.Rock:
				score = score + shapeScores.Rock + resultScores.Draw
			case player1Moves.Paper:
				score = score + shapeScores.Paper + resultScores.Draw
			case player1Moves.Scissors:
				score = score + shapeScores.Scissors + resultScores.Draw
			}
		}
	}

	return score
}
