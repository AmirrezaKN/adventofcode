package day04

import (
	"app/helper"
	"strconv"
	"strings"
	"testing"
)

func TestQuestion04_1(t *testing.T) {
	lines := helper.GetInputLines("./04_input.txt")
	pairs := []pair{}

	for _, line := range lines {
		p := strings.Split(line, ",")
		temp := pair{
			A: sectionRange{
				begin: func() int { t, _ := strconv.Atoi(strings.Split(p[0], "-")[0]); return t }(),
				end:   func() int { t, _ := strconv.Atoi(strings.Split(p[0], "-")[1]); return t }(),
			},
			B: sectionRange{
				begin: func() int { t, _ := strconv.Atoi(strings.Split(p[1], "-")[0]); return t }(),
				end:   func() int { t, _ := strconv.Atoi(strings.Split(p[1], "-")[1]); return t }(),
			},
		}
		pairs = append(pairs, temp)
	}

	type args struct {
		input []pair
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				input: pairs,
			},
			want: 644,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Question03_1(tt.args.input); got != tt.want {
				t.Errorf("Question03_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuestion04_2(t *testing.T) {
	lines := helper.GetInputLines("./04_input.txt")
	pairs := []pair{}

	for _, line := range lines {
		p := strings.Split(line, ",")
		temp := pair{
			A: sectionRange{
				begin: func() int { t, _ := strconv.Atoi(strings.Split(p[0], "-")[0]); return t }(),
				end:   func() int { t, _ := strconv.Atoi(strings.Split(p[0], "-")[1]); return t }(),
			},
			B: sectionRange{
				begin: func() int { t, _ := strconv.Atoi(strings.Split(p[1], "-")[0]); return t }(),
				end:   func() int { t, _ := strconv.Atoi(strings.Split(p[1], "-")[1]); return t }(),
			},
		}
		pairs = append(pairs, temp)
	}

	type args struct {
		input []pair
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				input: pairs,
			},
			want: 926,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Question03_2(tt.args.input); got != tt.want {
				t.Errorf("Question03_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
