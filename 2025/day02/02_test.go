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
	lols := strings.Split(str, ",")

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
					"11-22",
					"95-115",
					"998-1012",
					"1188511880-1188511890",
					"222220-222224",
					"1698522-1698528",
					"446443-446449",
					"38593856-38593862",
				},
			},
			want: 1227775554,
		},
		{
			name: "2",
			args: args{
				input: lols,
			},
			want: 1135,
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
	lols := strings.Split(str, ",")

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
					"11-22",
					"95-115",
					"998-1012",
					"1188511880-1188511890",
					"222220-222224",
					"1698522-1698528",
					"446443-446449",
					"38593856-38593862",
					"565653-565659",
					"824824821-824824827",
					"2121212118-2121212124",
				},
			},
			want: 4174379265,
		},
		{
			name: "2",
			args: args{
				input: lols,
			},
			want: 1135,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Question02_2(tt.args.input))
		})
	}
}
