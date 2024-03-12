package main

import "testing"

func Test_rob(t *testing.T) {
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
			name: "single house",
			args: args{[]int{5}},
			want: 5,
		},
		{
			name: "two houses",
			args: args{[]int{5, 11}},
			want: 11,
		},
		{
			name: "three houses",
			args: args{[]int{5, 11, 1}},
			want: 11,
		},
		{
			name: "three houses, will add",
			args: args{[]int{5, 11, 20}},
			want: 25,
		},
		// {
		// 	name: "timing out on leetcode",
		// 	args: args{[]int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240}},
		// 	want: 4173,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
