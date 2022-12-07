package day07

import (
	"app/helper"
	"testing"

	"gotest.tools/assert"
)

func TestQuestion07_1(t *testing.T) {
	sampleLines := helper.GetInputLines("07_sample_input.txt")
	sampleLines2 := helper.GetInputLines("07_sample_input2.txt")
	sampleLines3 := helper.GetInputLines("07_sample_input3.txt")
	lines := helper.GetInputLines("07_input.txt")
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sample",
			args: args{
				lines: sampleLines,
			},
			want: 95437,
		},
		{
			name: "sample3",
			args: args{
				sampleLines3,
			},
			want: 1390824,
		},
		{
			name: "sample2",
			args: args{
				sampleLines2,
			},
			want: 0,
		},
		{
			name: "1",
			args: args{
				lines: lines,
			},
			//1097622 is too low
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Question07_1(tt.args.lines))
		})
	}
}
