package main

import (
	"testing"
)

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "failing",
			args: args{
				nums: []int{-2, 3, -4},
			},
			want: 24,
		},
		{
			name: "failing 2",
			args: args{
				nums: []int{-2, 0, -1},
			},
			want: 0,
		},
		{
			name: "failing 3",
			args: args{
				nums: []int{3, -1, 4},
			},
			want: 4,
		},
		{
			name: "failing 4",
			args: args{
				nums: []int{2, -5, -2, -4, 3},
			},
			want: 24,
		},
		{
			name: "failing 5",
			args: args{
				nums: []int{0, -3, -2, -3, -2, 2, -3, 0, 1, -1},
			},
			want: 72,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
