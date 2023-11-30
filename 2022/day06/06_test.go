package day06

import (
	"app/helper"
	"testing"

	"gotest.tools/assert"
)

func TestQuestion06_1(t *testing.T) {
	lines := helper.GetInputLines("06_input.txt")
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{
				str: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			},
			want: 5,
		},
		{
			name: "02",
			args: args{
				str: "nppdvjthqldpwncqszvftbrmjlhg",
			},
			want: 6,
		},
		{
			name: "03",
			args: args{
				str: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			},
			want: 10,
		},
		{
			name: "04",
			args: args{
				str: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			},
			want: 11,
		},
		{
			name: "1",
			args: args{
				str: lines[0],
			},
			want: 1655,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Question06_1(tt.args.str))
		})
	}
}

func TestQuestion06_2(t *testing.T) {
	lines := helper.GetInputLines("06_input.txt")
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{
				str: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			},
			want: 19,
		},
		{
			name: "02",
			args: args{
				str: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			},
			want: 23,
		},
		{
			name: "03",
			args: args{
				str: "nppdvjthqldpwncqszvftbrmjlhg",
			},
			want: 23,
		},
		{
			name: "04",
			args: args{
				str: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			},
			want: 29,
		},
		{
			name: "05",
			args: args{
				str: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			},
			want: 26,
		},
		{
			name: "1",
			args: args{
				str: lines[0],
			},
			want: 2665,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Question06_2(tt.args.str))
		})
	}
}
