package day01

import (
	"io"
	"os"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestPart1(t *testing.T) {
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, 0777)
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
					"7 6 4 2 1",
					"1 2 7 8 9",
					"9 7 6 2 1",
					"1 3 2 4 5",
					"8 6 4 4 1",
					"1 3 6 7 9",
				},
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				input: lols,
			},
			want: 510,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Part1(tt.args.input))
		})
	}
}

func TestPart2(t *testing.T) {
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, 0777)
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
					"35 38 36 34",
					"70 72 74 77 78 83",
					"7 6 4 2 1",
					"1 2 7 8 9",
					"9 7 6 2 1",
					"1 3 2 4 5",
					"8 6 4 4 1",
					"1 3 6 7 9",
				},
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				input: lols,
			},
			want: 553,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Part2(tt.args.input))
		})
	}
}
