package day05

import (
	"app/helper"
	"strconv"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestQuestion05_1(t *testing.T) {
	lines := helper.GetInputLines("./05_input.txt")

	index := 0
	for i, line := range lines {
		if line == "" {
			index = i
			break
		}
	}

	baseStacks := lines[:index]
	movesStr := lines[index+1:]

	stacks := [9][]string{}
	for i := len(baseStacks) - 2; i >= 0; i-- {
		items := strings.Split(strings.ReplaceAll(baseStacks[i], "    ", " "), " ")

		for j, item := range items {
			item = strings.Trim(item, "[")
			item = strings.Trim(item, "]")
			if item != "" {
				stacks[j] = append(stacks[j], item)
			}
		}
	}

	moves := []move{}
	for _, item := range movesStr {
		moveStrSlice := strings.Split(item, " ")
		moves = append(moves, move{
			amount:      func() int { t, _ := strconv.Atoi(moveStrSlice[1]); return t }(),
			source:      func() int { t, _ := strconv.Atoi(moveStrSlice[3]); return t }(),
			destination: func() int { t, _ := strconv.Atoi(moveStrSlice[5]); return t }(),
		})
	}

	type args struct {
		stacks [9][]string
		moves  []move
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				stacks: stacks,
				moves:  moves,
			},
			want: "SPFMVDTZT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Question05_1(tt.args.stacks, tt.args.moves))
		})
	}
}

func TestQuestion05_2(t *testing.T) {
	lines := helper.GetInputLines("./05_input.txt")

	index := 0
	for i, line := range lines {
		if line == "" {
			index = i
			break
		}
	}

	baseStacks := lines[:index]
	movesStr := lines[index+1:]

	stacks := [9][]string{}
	for i := len(baseStacks) - 2; i >= 0; i-- {
		items := strings.Split(strings.ReplaceAll(baseStacks[i], "    ", " "), " ")

		for j, item := range items {
			item = strings.Trim(item, "[")
			item = strings.Trim(item, "]")
			if item != "" {
				stacks[j] = append(stacks[j], item)
			}
		}
	}

	moves := []move{}
	for _, item := range movesStr {
		moveStrSlice := strings.Split(item, " ")
		moves = append(moves, move{
			amount:      func() int { t, _ := strconv.Atoi(moveStrSlice[1]); return t }(),
			source:      func() int { t, _ := strconv.Atoi(moveStrSlice[3]); return t }(),
			destination: func() int { t, _ := strconv.Atoi(moveStrSlice[5]); return t }(),
		})
	}

	type args struct {
		stacks [9][]string
		moves  []move
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				stacks: stacks,
				moves:  moves,
			},
			want: "ZFSJBPRFP",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Question05_2(tt.args.stacks, tt.args.moves))
		})
	}
}
