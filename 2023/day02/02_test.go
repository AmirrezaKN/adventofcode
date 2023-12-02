package day01

import (
	"io"
	"os"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestQuestion021(t *testing.T) {
	f, err := os.OpenFile("./02_input.txt", os.O_RDONLY, 0777)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	str := string(bytes)
	lols := strings.Split(str, "\n")

	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				input: []string{
					"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
					"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
					"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
					"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
					"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
				},
			},
			want: 8,
		},
		{
			name: "2",
			args: args{
				input: lols,
			},
			want: 2551,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Question02_1(tt.args.input))
		})
	}
}

func TestQuestion022(t *testing.T) {
	f, err := os.OpenFile("./02_input.txt", os.O_RDONLY, 0777)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	str := string(bytes)
	lols := strings.Split(str, "\n")

	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				input: []string{
					"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
					"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
					"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
					"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
					"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
				},
			},
			want: 2286,
		},
		{
			name: "2",
			args: args{
				input: lols,
			},
			want: 62811,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Question02_2(tt.args.input))
		})
	}
}
