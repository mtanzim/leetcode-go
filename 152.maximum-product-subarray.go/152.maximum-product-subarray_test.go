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
			want: [][]int{{}},
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
		// {
		// 	name: "base - triple",
		// 	args: args{
		// 		[]int{-2, 3, -4},
		// 	},
		// 	want: [][]int{{1}, {1, 2}, {2}},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarrays(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
				nums: []int{-2,3,-4},
			},
			want: 24,
		},
		{
			name: "failing",
			args: args{
				nums: []int{-2,0,-1},
			},
			want: 0,
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
