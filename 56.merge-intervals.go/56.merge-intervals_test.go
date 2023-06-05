package main

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "given 1",
			args: args{
				intervals: [][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}},
			},
			want: [][]int{[]int{1, 6}, []int{8, 10}, []int{15, 18}},
		},
		{
			name: "given edge case",
			args: args{
				intervals: [][]int{[]int{1, 4}, []int{0, 4}},
			},
			want: [][]int{[]int{0, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
