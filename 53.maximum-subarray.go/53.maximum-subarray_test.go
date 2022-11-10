package main

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "single elem",
			args: args{[]int{5}},
			want: 5,
		},
		{
			name: "two elem, adds up to max",
			args: args{[]int{5, 2}},
			want: 7,
		},
		{
			name: "two elems, sum is smaller than idx 0",
			args: args{[]int{5, -2}},
			want: 5,
		},
		{
			name: "two elems, sum is smaller than idx 1",
			args: args{[]int{-5, 6}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
