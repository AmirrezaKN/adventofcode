package day03

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
					"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
				},
			},
			want: 161,
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
					"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
				},
			},
			want: 48,
		},
		{
			name: "2",
			args: args{
				input: []string{str},
			},
			want: 92082041,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Part2(tt.args.input))
		})
	}
}
