package day01

import (
	"io"
	"os"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestQuestion031(t *testing.T) {
	f, err := os.OpenFile("./03_input.txt", os.O_RDONLY, 0777)
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
					"987654321111111",
					"811111111111119",
					"234234234234278",
					"818181911112111",
				},
			},
			want: 357,
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
			assert.Equal(t, tt.want, Question03_1(tt.args.input))
		})
	}
}

func TestQuestion032(t *testing.T) {
	f, err := os.OpenFile("./03_input.txt", os.O_RDONLY, 0777)
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
					"987654321111111",
					"811111111111119",
					"234234234234278",
					"818181911112111",
				},
			},
			want: 3121910778619,
		},
		{
			name: "2",
			args: args{
				input: lols,
			},
			want: 168027167146027,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Question03_2(tt.args.input))
		})
	}
}
