package day03

import (
	"app/helper"
	"testing"
)

func TestQuestion02_1(t *testing.T) {
	lines := helper.GetInputLines("./03_input.txt")

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
				input: lines,
			},
			want: 7889,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Question02_1(tt.args.input); got != tt.want {
				t.Errorf("Question02_1() = %v, want %v", got, tt.want)
			}
		})
	}
}
