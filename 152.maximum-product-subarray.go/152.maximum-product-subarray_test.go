package main

import (
	"reflect"
	"testing"
)

func Test_subarrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "base - empty",
			args: args{
				[]int{},
			},
			want: [][]int{},
		},
		{
			name: "base - single",
			args: args{
				[]int{0},
			},
			want: [][]int{{0}},
		},
		{
			name: "base - double",
			args: args{
				[]int{1, 2},
			},
			want: [][]int{{1}, {1, 2}, {2}},
		},
		{
			name: "base - triple",
			args: args{
				[]int{1, 2, 3},
			},
			want: [][]int{{1}, {1, 2}, {2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarrays(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
