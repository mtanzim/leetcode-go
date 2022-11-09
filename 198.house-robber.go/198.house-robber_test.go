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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
