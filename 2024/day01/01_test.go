package day01

import (
	"io"
	"os"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestQuestion011(t *testing.T) {
	f, err := os.OpenFile("./01_input.txt", os.O_RDONLY, 0777)
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
					"3   4",
					"4   3",
					"2   5",
					"1   3",
					"3   9",
					"3   3",
				},
			},
			want: 11,
		},
		{
			name: "2",
			args: args{
				input: lols,
			},
			want: 2057374,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Question01_1(tt.args.input))
		})
	}
}

func TestQuestion012(t *testing.T) {
	f, err := os.OpenFile("./01_input.txt", os.O_RDONLY, 0777)
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
					"3   4",
					"4   3",
					"2   5",
					"1   3",
					"3   9",
					"3   3",
				},
			},
			want: 31,
		},
		{
			name: "2",
			args: args{
				input: lols,
			},
			want: 54076,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Question01_2(tt.args.input))
		})
	}
}
