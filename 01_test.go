package main

import (
	"io/ioutil"
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

	bytes, err := ioutil.ReadAll(f)
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
				input: lols,
			},
			want: 69836,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Question011(tt.args.input))
		})
	}
}

func TestQuestion012(t *testing.T) {
	f, err := os.OpenFile("./01_input.txt", os.O_RDONLY, 0777)
	if err != nil {
		t.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(f)
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
				input: lols,
			},
			want: 207968,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Question012(tt.args.input))
		})
	}
}
