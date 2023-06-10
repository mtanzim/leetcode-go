package main

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "base",
			args: args{intervals: [][]int{[]int{1, 3}, []int{6, 9}}, newInterval: []int{2, 5}},
			want: [][]int{[]int{1,5}, []int{6,9}},
		},
		{
			name: "base 2",
			args: args{intervals: [][]int{[]int{1, 2}, []int{3, 5}, []int{6, 7}, []int{8, 10}, []int{12, 16}}, newInterval: []int{4, 8}},
			want: [][]int{[]int{1,2}, []int{3,10}, []int{12,16}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
