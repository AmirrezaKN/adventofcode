package main

import (
	"os"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestQuestion02_1(t *testing.T) {
	inputBytes, err := os.ReadFile("./02_input.txt")
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(string(inputBytes), "\n")
	argsInput := make([]q2round, 0)
	for _, line := range lines {
		moves := strings.Split(line, " ")

		argsInput = append(argsInput, q2round{
			P1: moves[0],
			P2: moves[1],
		})
	}

	type args struct {
		input []q2round
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				input: argsInput,
			},
			want: 13565,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Question02_1(tt.args.input))
		})
	}
}

func TestQuestion02_2(t *testing.T) {
	inputBytes, err := os.ReadFile("./02_input.txt")
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(string(inputBytes), "\n")
	argsInput := make([]q2round, 0)
	for _, line := range lines {
		moves := strings.Split(line, " ")

		argsInput = append(argsInput, q2round{
			P1: moves[0],
			P2: moves[1],
		})
	}

	type args struct {
		input []q2round
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				input: argsInput,
			},
			want: 12424,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Question02_2(tt.args.input))
		})
	}
}
