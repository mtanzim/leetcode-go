package main

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "base",
			args: args{[]int{1}},
			want: [][]int{[]int{1}},
		},
		{
			name: "base with len 2",
			args: args{[]int{0, 1}},
			want: [][]int{[]int{0, 1}, []int{1, 0}},
		},
		{
			name: "base with len 3",
			args: args{[]int{1, 2, 3}},
			want: [][]int{[]int{1, 2, 3}, []int{1, 3, 2}, []int{2, 1, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
