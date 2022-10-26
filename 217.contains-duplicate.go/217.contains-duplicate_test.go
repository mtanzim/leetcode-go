package main

import "testing"

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "basic",
			args: args{[]int{1, 2, 3, 1}},
			want: true,
		},
		{name: "basic2",
			args: args{[]int{1, 2, 3, 4}},
			want: false,
		},
		{name: "basic3",
			args: args{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
