package day05

import (
	"io"
	"os"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestQuestion051(t *testing.T) {
	f, err := os.OpenFile("./05_input.txt", os.O_RDONLY, 0777)
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
					"3-5",
					"10-14",
					"16-20",
					"12-18",
					"",
					"1",
					"5",
					"8",
					"11",
					"17",
					"32",
				},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				input: lols,
			},
			want: 16973,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Question05_1(tt.args.input))
		})
	}
}

func TestQuestion052(t *testing.T) {
	f, err := os.OpenFile("./05_input.txt", os.O_RDONLY, 0777)
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
		want uint
	}{
		{
			name: "1",
			args: args{
				input: []string{
					"3-5",
					"10-14",
					"16-20",
					"12-18",
					"",
					"1",
					"5",
					"8",
					"11",
					"17",
					"32",
				},
			},
			want: 14,
		},
		{
			name: "2",
			args: args{
				input: lols,
			},
			want: 336790092076620,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Question05_2(tt.args.input))
		})
	}
}
