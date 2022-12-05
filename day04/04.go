package day04

type sectionRange struct {
	begin int
	end   int
}
type pair struct {
	A, B sectionRange
}

func Question04_1(input []pair) int {
	total := 0

	for _, p := range input {
		if (p.A.begin >= p.B.begin) && (p.A.end <= p.B.end) {
			total++
		} else if (p.B.begin >= p.A.begin) && (p.B.end <= p.A.end) {
			total++
		}
	}

	return total
}

func Question04_2(input []pair) int {
	total := 0

	for _, p := range input {
		if (p.A.begin >= p.B.begin) && (p.A.end <= p.B.end) {
			total++
		} else if (p.B.begin >= p.A.begin) && (p.B.end <= p.A.end) {
			total++
		} else if (p.A.begin <= p.B.end) && (p.A.end >= p.B.begin) {
			total++
		}
	}

	return total
}
